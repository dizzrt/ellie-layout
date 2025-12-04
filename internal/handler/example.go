package handler

import (
	"context"

	"github.com/dizzrt/ellie-layout/api/gen/example"
	"github.com/dizzrt/ellie-layout/internal/application"
)

type ExampleHandler struct {
	example.UnimplementedExampleServiceServer

	exampleApp application.ExampleApplication
}

func NewExampleHandler(exampleApp application.ExampleApplication) *ExampleHandler {
	return &ExampleHandler{
		exampleApp: exampleApp,
	}
}

func (h *ExampleHandler) Hello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
	return &example.HelloResponse{
		Message: "hello " + req.Name,
	}, nil
}
