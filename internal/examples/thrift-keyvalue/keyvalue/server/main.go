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

package main

import (
	"context"
	"fmt"
	"log"
	gohttp "net/http"
	"sync"

	"go.uber.org/yarpc/internal/examples/thrift-keyvalue/keyvalue/kv"
	"go.uber.org/yarpc/internal/examples/thrift-keyvalue/keyvalue/kv/keyvalueserver"
	"go.uber.org/yarpc/x/yarpcmeta"

	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/http"
	"go.uber.org/yarpc/transport/tchannel"
)

type handler struct {
	sync.RWMutex

	items map[string]string
}

func (h *handler) GetValue(ctx context.Context, key *string) (string, error) {
	h.RLock()
	defer h.RUnlock()

	if value, ok := h.items[*key]; ok {
		return value, nil
	}

	return "", &kv.ResourceDoesNotExist{Key: *key}
}

func (h *handler) SetValue(ctx context.Context, key *string, value *string) error {
	h.Lock()
	h.items[*key] = *value
	h.Unlock()
	return nil
}

func main() {
	go func() {
		if err := gohttp.ListenAndServe(":3242", nil); err != nil {
			log.Fatal(err)
		}
	}()
	tchannelTransport, err := tchannel.NewChannelTransport(
		tchannel.ServiceName("keyvalue"),
		tchannel.ListenAddr(":28941"),
	)
	if err != nil {
		log.Fatal(err)
	}

	httpTransport := http.NewTransport()
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: "keyvalue",
		Inbounds: yarpc.Inbounds{
			tchannelTransport.NewInbound(),
			httpTransport.NewInbound(":24034"),
		},
	})

	handler := handler{items: make(map[string]string)}
	dispatcher.Register(keyvalueserver.New(&handler))

	yarpcmeta.Register(dispatcher)

	if err := dispatcher.Start(); err != nil {
		fmt.Println("error:", err.Error())
	}

	select {} // block forever
}
