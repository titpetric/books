package kinesis

import (
	"time"

	"github.com/apex/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	k "github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/jpillora/backoff"
)

const (
	maxRecordsPerRequest = 500
)

type Config struct {
	// StreamName is the Kinesis stream.
	StreamName string

	// FlushInterval is a regular interval for flushing the buffer. Defaults to 1s.
	FlushInterval time.Duration

	// BufferSize determines the batch request size. Must not exceed 500. Defaults to 500.
	BufferSize int

	// BacklogSize determines the channel capacity before Put() will begin blocking. Defaults to 500.
	BacklogSize int

	// Backoff determines the backoff strategy for record failures.
	Backoff backoff.Backoff

	// Logger is the logger used. Defaults to log.Log.
	Logger log.Interface

	// Client is the Kinesis API implementation.
	Client kinesisiface.KinesisAPI
}

// defaults for configuration.
func (c *Config) defaults() {
	if c.Client == nil {
		c.Client = k.New(session.New(aws.NewConfig()))
	}

	if c.Logger == nil {
		c.Logger = log.Log
	}

	c.Logger = c.Logger.WithFields(log.Fields{
		"package": "kinesis",
	})

	if c.StreamName == "" {
		c.Logger.Fatal("StreamName required")
	}

	c.Logger = c.Logger.WithFields(log.Fields{
		"stream": c.StreamName,
	})

	if c.BufferSize == 0 {
		c.BufferSize = maxRecordsPerRequest
	}

	if c.BufferSize > maxRecordsPerRequest {
		c.Logger.Fatal("BufferSize exceeds 500")
	}

	if c.BacklogSize == 0 {
		c.BacklogSize = maxRecordsPerRequest
	}

	if c.FlushInterval == 0 {
		c.FlushInterval = time.Second
	}
}
