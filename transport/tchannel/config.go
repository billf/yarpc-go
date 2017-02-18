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

package tchannel

import (
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/x/config"
)

// TransportConfig TODO
type TransportConfig struct {
	Address string `config:"address"`
	Service string `config:"service"`
}

// BuildTransport TODO
func (t *TransportConfig) BuildTransport() (transport.Transport, error) {
	var opts []TransportOption
	if t.Address != "" {
		opts = append(opts, ListenAddr(t.Address))
	}
	if t.Service != "" {
		opts = append(opts, ServiceName(t.Service))
	}
	return NewTransport(opts...)
}

// InboundConfig TODO
type InboundConfig struct{}

// BuildInbound TODO
func (c *InboundConfig) BuildInbound(t transport.Transport) (transport.Inbound, error) {
	return t.(*Transport).NewInbound(), nil
}

// OutboundConfig TODO
type OutboundConfig struct {
	Address string `config:"address"`
}

// BuildUnaryOutbound TODO
func (c *OutboundConfig) BuildUnaryOutbound(t transport.Transport) (transport.UnaryOutbound, error) {
	return t.(*Transport).NewSingleOutbound(c.Address), nil
}

// RegisterTransport TODO
func RegisterTransport(l *config.Loader) error {
	return l.RegisterTransport(config.TransportSpec{
		Name:                    "tchannel",
		TransportConfigType:     &TransportConfig{},
		InboundConfigType:       &InboundConfig{},
		UnaryOutboundConfigType: &OutboundConfig{},
	})
}
