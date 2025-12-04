package domain

import (
	example_biz "github.com/dizzrt/ellie-layout/internal/domain/example/biz"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	example_biz.NewExampleBiz,
)
