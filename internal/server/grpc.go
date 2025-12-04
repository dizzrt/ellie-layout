package server

import (
	"github.com/dizzrt/ellie-layout/api/gen/example"
	"github.com/dizzrt/ellie-layout/internal/conf"
	"github.com/dizzrt/ellie-layout/internal/handler"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/transport/grpc"
)

func NewGRPCServer(c *conf.AppConfig, logger log.LogWriter, exampleHandler *handler.ExampleHandler) *grpc.Server {
	opts := []grpc.ServerOption{}

	grpcServerConf := c.Server.GRPC
	if grpcServerConf.Addr != "" {
		opts = append(opts, grpc.Address(grpcServerConf.Addr))
	}

	srv := grpc.NewServer(opts...)
	example.RegisterExampleServiceServer(srv, exampleHandler)

	return srv
}
