package biz

import "context"

var _ ExampleBiz = (*exampleBiz)(nil)

type ExampleBiz interface {
	Hello(ctx context.Context, name string) (string, error)
}

type exampleBiz struct {
}

func NewExampleBiz() ExampleBiz {
	return &exampleBiz{}
}

func (biz *exampleBiz) Hello(ctx context.Context, name string) (string, error) {
	return "hello " + name, nil
}
