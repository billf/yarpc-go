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
	"io/ioutil"
	"log"

	"github.com/kr/pretty"

	"go.uber.org/yarpc"
	"go.uber.org/yarpc/encoding/json"
)

func main() {
	config, err := ioutil.ReadFile("server.yaml")
	if err != nil {
		log.Fatal(err)
	}

	builder, err := loader.LoadYAML(config)
	if err != nil {
		log.Fatal(err)
	}
	pretty.Println("Loaded configuration:", builder)

	cfg, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	d := yarpc.NewDispatcher(cfg)
	d.Register(json.Procedure("echo", echo))
	d.Register(json.OnewayProcedure("enqueue", enqueue))

	if err := d.Start(); err != nil {
		log.Fatal(err)
	}
	defer d.Stop()

	select {}
}

func echo(ctx context.Context, req map[string]string) (map[string]string, error) {
	return req, nil
}

func enqueue(ctx context.Context, req map[string]string) error {
	fmt.Println("received oneway request", req)
	return nil
}
