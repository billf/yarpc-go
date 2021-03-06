// Code generated by thriftrw-plugin-yarpc
// @generated

package emptyserviceclient

import (
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
)

// Interface is a client for the EmptyService service.
type Interface interface {
}

// New builds a new client for the EmptyService service.
//
// 	client := emptyserviceclient.New(dispatcher.ClientConfig("emptyservice"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "EmptyService",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(func(c transport.ClientConfig) Interface {
		return New(c)
	})
}

type client struct {
	c thrift.Client
}
