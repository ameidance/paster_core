// Code generated by Kitex v0.0.1. DO NOT EDIT.

package pastercoreservice

import (
	"github.com/ameidance/paster_core/model/dto/kitex_gen/ameidance/paster/core"
	"github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler core.PasterCoreService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
