package golf

import (
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

// Compression type to use for GELF messages that are sent
const (
	COMP_NONE = iota // No compression
	COMP_GZIP        // gzip compression
	COMP_ZLIB        // zlib compression
)

type Client struct {
	hostname string

	conn net.Conn

	queue      []*Message
	queueMutex sync.Mutex

	msgChan  chan *Message
	queueCtl chan int
	sendCtl  chan int

	config ClientConfig
}

// Configuration used when creating a server instance
type ClientConfig struct {
	ChunkSize   int // The data size for each chunk sent to the server
	Compression int // Compression to use for messagec.
}

/*
Create a new Client instance with the default values for ClientConfig:

 {
	ChunkSize: 1420,
	Compression: COMP_GZIP,
 }
*/
func NewClient() (*Client, error) {
	cc := ClientConfig{
		ChunkSize:   1420,
		Compression: COMP_GZIP,
	}
	return NewClientWithConfig(cc)
}

// Create a new Client instance with the given ClientConfig
func NewClientWithConfig(config ClientConfig) (*Client, error) {
	c := &Client{
		config: config,
		queue:  make([]*Message, 0),

		msgChan:  make(chan *Message, 500),
		queueCtl: make(chan int),
		sendCtl:  make(chan int),
	}

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	c.hostname = host

	return c, nil
}

// Connect to a GELF server at the given URI.
func (c *Client) Dial(uri string) error {
	parsedUri, err := url.Parse(uri)
	if err != nil {
		return err
	}

	if !strings.Contains(parsedUri.Host, ":") {
		parsedUri.Host = parsedUri.Host + ":12201"
	}

	switch parsedUri.Scheme {
	case "udp":
	case "tcp":
	default:
		return errors.New("Unsupported scheme provided")
	}

	conn, err := net.Dial(parsedUri.Scheme, parsedUri.Host)
	if err != nil {
		return err
	}
	c.conn = conn

	go c.queueReceiver()
	go c.msgSender()

	return nil
}

// Close the connection to the server. This call will block until all the
// currently queued messages for the client are sent.
func (c *Client) Close() error {
	if c.conn == nil {
		// Already shut down so it doesn't need to run again
		return nil
	}

	// First quit the queue and wait for it to respond
	// that it's quit
	c.queueCtl <- 1
	for {
		quitVal := <-c.queueCtl
		if quitVal == 2 {
			break
		}
		c.queueCtl <- quitVal
	}

	// Then quit the sender and wait for it to respond
	// that it's quit
	c.sendCtl <- 1
	for {
		quitVal := <-c.sendCtl
		if quitVal == 2 {
			break
		}
		c.sendCtl <- quitVal
	}

	err := c.conn.Close()
	if err != nil {
		return err
	}
	c.conn = nil

	return nil
}

// Queue the given message at the end of the message queue
func (c *Client) QueueMsg(msg *Message) error {
	if msg.Timestamp == nil {
		curTime := time.Now()
		msg.Timestamp = &curTime
	}

	c.msgChan <- msg
	return nil
}

func (c *Client) queueReceiver() {
	for {
		select {
		case msg := <-c.msgChan:
			c.queueMutex.Lock()
			c.queue = append(c.queue, msg)
			c.queueMutex.Unlock()
		case quitVal := <-c.queueCtl:
			if quitVal == 1 {
				// Don't quit if there are still
				// messages in the channel
				if len(c.msgChan) > 0 {
					c.queueCtl <- 1
					continue
				}
				c.queueCtl <- 2
				return
			}
		}
	}
}

func (c *Client) msgSender() {
	var msg *Message
	for {
		if len(c.queue) > 0 {
			c.queueMutex.Lock()
			msg, c.queue = c.queue[0], c.queue[1:]
			c.queueMutex.Unlock()

			data, err := generateMsgJson(msg)
			if err != nil {
				// TODO Not sure what to do at this point? Fail the
				// message silently?
				// Might be able to add an error channel that the
				// user can watch for errors
				continue
			}
			err = c.writeMsg(data, c.conn, COMP_GZIP)
			if err != nil {
				// TODO Same as above...
			}
		} else {
			time.Sleep(1 * time.Second)

			select {
			case quitVal := <-c.sendCtl:
				if quitVal == 1 {
					if len(c.queue) > 0 {
						c.sendCtl <- 1
						continue
					}
					c.sendCtl <- 2
					return
				}
			default:
			}
		}
	}
}

func (c *Client) writeMsg(data string, w io.Writer, compression int) error {
	chnk, err := newChunker(w, c.config.ChunkSize)
	if err != nil {
		return err
	}
	defer chnk.Flush()

	switch compression {
	case COMP_GZIP:
		gz, err := gzip.NewWriterLevel(chnk, gzip.DefaultCompression)
		if err != nil {
			return err
		}
		gz.Write([]byte(data))
		gz.Close()
	case COMP_ZLIB:
		zz, err := zlib.NewWriterLevel(chnk, zlib.DefaultCompression)
		if err != nil {
			return err
		}
		zz.Write([]byte(data))
		zz.Close()
	default:
		chnk.Write([]byte(data))
	}

	return nil
}
