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

package http

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/x/config"
)

// TransportConfig TODO
type TransportConfig struct {
	KeepAlive time.Duration `config:"keepAlive"`
}

// BuildTransport TODO
func (tc *TransportConfig) BuildTransport() (transport.Transport, error) {
	var opts []TransportOption
	if tc.KeepAlive > 0 {
		opts = append(opts, KeepAlive(tc.KeepAlive))
	}

	return NewTransport(opts...), nil
}

// InboundConfig TODO
type InboundConfig struct {
	Address string `config:"address"`
}

// BuildInbound TODO
func (ic *InboundConfig) BuildInbound(t transport.Transport) (transport.Inbound, error) {
	if ic.Address == "" {
		return nil, fmt.Errorf("inbound address is required")
	}

	return t.(*Transport).NewInbound(ic.Address), nil
}

// OutboundConfig TODO
type OutboundConfig struct {
	Host   string `config:"host"`
	Port   int    `config:"port"`
	Path   string `config:"path"`
	Scheme string `config:"scheme"`
}

// BuildUnaryOutbound TODO
func (oc *OutboundConfig) BuildUnaryOutbound(t transport.Transport) (transport.UnaryOutbound, error) {
	return oc.buildOutbound(t)
}

// BuildOnewayOutbound TODO
func (oc *OutboundConfig) BuildOnewayOutbound(t transport.Transport) (transport.OnewayOutbound, error) {
	return oc.buildOutbound(t)
}

func (oc *OutboundConfig) buildOutbound(t transport.Transport) (*Outbound, error) {
	if oc.Host == "" {
		return nil, errors.New("outbound host is required")
	}
	if oc.Port == 0 {
		return nil, errors.New("outbound port is required")
	}

	if oc.Scheme == "" {
		oc.Scheme = "http"
	}

	if oc.Path != "" && oc.Path[0] == '/' {
		oc.Path = oc.Path[1:]
	}

	url := fmt.Sprintf("%s://%s:%d/%s", oc.Scheme, oc.Host, oc.Port, oc.Path)
	return t.(*Transport).NewSingleOutbound(url), nil
}

// TODO: presets

// RegisterTransport TODO
func RegisterTransport(l *config.Loader) error {
	return l.RegisterTransport(config.TransportSpec{
		Name:                     "http",
		TransportConfigType:      &TransportConfig{},
		InboundConfigType:        &InboundConfig{},
		UnaryOutboundConfigType:  &OutboundConfig{},
		OnewayOutboundConfigType: &OutboundConfig{},
	})
}
