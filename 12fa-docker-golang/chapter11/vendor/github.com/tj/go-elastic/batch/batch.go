// Package batch lets you buffer bulk documents for insert. None of the methods
// provided are thread-safe, you must synchronize if desired.
package batch

import (
	"bytes"
	"encoding/json"
	"io"
)

// Elasticsearch interface.
type Elasticsearch interface {
	Bulk(io.Reader) error
}

// Index metadata.
type Index struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	Routing string `json:"_routing,omitempty"`
	ID      string `json:"_id,omitempty"`
}

// IndexOp is an index operation.
type IndexOp struct {
	Index Index `json:"index"`
}

// Batch indexes docs in bulk for reporting. Currently documents
// are flushed in a single write, however may allow streaming
// in the future.
type Batch struct {
	Elastic Elasticsearch // Elasticsearch implementation
	Docs    []interface{} // Docs buffered
	Index   string        // Index name
	Type    string        // Type name
}

// Add document.
func (b *Batch) Add(doc interface{}) {
	b.Docs = append(b.Docs, doc)
}

// Size returns the number of documents pending flush.
func (b *Batch) Size() int {
	return len(b.Docs)
}

// Bytes returns the request body.
func (b *Batch) Bytes() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)

	op := IndexOp{
		Index: Index{
			Index: b.Index,
			Type:  b.Type,
		},
	}

	for _, doc := range b.Docs {
		if err := enc.Encode(op); err != nil {
			return nil, err
		}

		if err := enc.Encode(doc); err != nil {
			return nil, err
		}
	}

	return buf, nil
}

// Flush checks in bulk.
func (b *Batch) Flush() (err error) {
	if b.Size() == 0 {
		return nil
	}

	buf, err := b.Bytes()
	if err != nil {
		return err
	}

	b.Docs = nil

	return b.Elastic.Bulk(buf)
}
