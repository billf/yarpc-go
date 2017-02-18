// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package redis

import (
	"errors"
	"time"

	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/x/config"
)

// TransportConfig TODO
type TransportConfig struct {
	Address string `config:"address"`
}

// BuildTransport TODO
func (c *TransportConfig) BuildTransport() (transport.Transport, error) {
	if c.Address == "" {
		return nil, errors.New("address is required")
	}

	return NewRedis5Client(c.Address), nil
}

// OutboundConfig TODO
type OutboundConfig struct {
	QueueKey string `config:"queueKey"`
}

// BuildOnewayOutbound TODO
func (c *OutboundConfig) BuildOnewayOutbound(t transport.Transport) (transport.OnewayOutbound, error) {
	if c.QueueKey == "" {
		return nil, errors.New("queue key is required")
	}

	return NewOnewayOutbound(t.(Client), c.QueueKey), nil
}

// InboundConfig TODO
type InboundConfig struct {
	QueueKey      string        `config:"queueKey"`
	ProcessingKey string        `config:"processingKey"`
	Timeout       time.Duration `config:"timeout"`
}

// BuildInbound TODO
func (c *InboundConfig) BuildInbound(t transport.Transport) (transport.Inbound, error) {
	if c.QueueKey == "" {
		return nil, errors.New("queue key is required")
	}

	if c.ProcessingKey == "" {
		return nil, errors.New("processing key is required")
	}

	if c.Timeout == 0 {
		c.Timeout = time.Second
	}

	return NewInbound(t.(Client), c.QueueKey, c.ProcessingKey, c.Timeout), nil
}

// RegisterTransport TODO
func RegisterTransport(l *config.Loader) error {
	return l.RegisterTransport(config.TransportSpec{
		Name:                     "redis",
		TransportConfigType:      &TransportConfig{},
		InboundConfigType:        &InboundConfig{},
		OnewayOutboundConfigType: &OutboundConfig{},
	})
}
