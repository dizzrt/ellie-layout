package infra

import (
	"github.com/dizzrt/ellie-layout/internal/infra/foundation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	foundation.ProviderSet,
)
