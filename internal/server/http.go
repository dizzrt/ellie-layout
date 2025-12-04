package server

import (
	"github.com/dizzrt/ellie-layout/api/gen/example"
	"github.com/dizzrt/ellie-layout/internal/conf"
	"github.com/dizzrt/ellie-layout/internal/handler"
	"github.com/dizzrt/ellie/log"
	"github.com/dizzrt/ellie/transport/http"
)

func NewHTTPServer(c *conf.AppConfig, logger log.LogWriter, exampleHandler *handler.ExampleHandler) *http.Server {
	opts := []http.ServerOption{}

	httpServerConf := c.Server.HTTP

	if httpServerConf.Addr != "" {
		opts = append(opts, http.Address(httpServerConf.Addr))
	}

	srv := http.NewServer(opts...)
	example.RegisterExampleServiceHTTPServer(srv, exampleHandler)

	return srv
}
