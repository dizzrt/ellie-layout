package application

import (
	"context"

	"github.com/dizzrt/ellie-layout/api/gen/example"
	"github.com/dizzrt/ellie-layout/internal/domain/example/biz"
)

var _ ExampleApplication = (*exampleApplication)(nil)

type ExampleApplication interface {
	Hello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error)
}

type exampleApplication struct {
	exampleBiz biz.ExampleBiz
}

func NewExampleApplication(exampleBiz biz.ExampleBiz) ExampleApplication {
	return &exampleApplication{
		exampleBiz: exampleBiz,
	}
}

func (app *exampleApplication) Hello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
	msg, err := app.exampleBiz.Hello(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	return &example.HelloResponse{
		Message: msg,
	}, nil
}
