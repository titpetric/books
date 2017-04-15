package golf

import (
	"errors"
	"github.com/satori/go.uuid"
	"io"
	"math"
)

type chunker struct {
	chunkSize int
	buff      []byte
	w         io.Writer
}

func newChunker(w io.Writer, chunkSize int) (*chunker, error) {
	if chunkSize < 13 {
		return nil, errors.New("Chunk size must be at least 13.")
	}

	c := &chunker{
		chunkSize: chunkSize,
		buff:      make([]byte, 0),
		w:         w,
	}

	return c, nil
}

func (c *chunker) reset() {
	c.buff = make([]byte, 0)
}
func (c *chunker) Write(p []byte) (int, error) {
	c.buff = append(c.buff, p...)
	return len(p), nil
}

func (c *chunker) Flush() error {
	idFull := uuid.NewV4()
	err := c.flushWithId(idFull.Bytes()[0:8])
	return err
}

func (c *chunker) flushWithId(id []byte) error {
	if len(id) < 8 || len(id) > 8 {
		return errors.New("id length must be equal to 8")
	}

	offset := 0
	buffLen := len(c.buff)
	chunkSize := c.chunkSize - 12

	// Reusing this buffer may cause problems with duplicate data being sent
	// if the data isn't written to something else by the io.Writer before
	// the chunk's data is updated.
	chunkBuff := make([]byte, c.chunkSize)
	copy(chunkBuff[0:2], []byte{0x1e, 0x0f})
	copy(chunkBuff[2:10], id)

	totalChunks := int(math.Ceil(float64(buffLen) / float64(chunkSize)))
	chunkBuff[11] = byte(totalChunks)

	for {
		left := buffLen - offset
		if left > chunkSize {
			copy(chunkBuff[12:], c.buff[offset:offset+chunkSize])
			c.w.Write(chunkBuff)
		} else {
			copy(chunkBuff[12:], c.buff[offset:offset+left])
			c.w.Write(chunkBuff[0 : left+12])
			break
		}

		offset += chunkSize
		chunkBuff[10] += 1
	}

	c.reset()
	return nil
}
