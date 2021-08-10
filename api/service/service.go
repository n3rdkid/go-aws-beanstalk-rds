package service

import (
	"go.uber.org/fx"
)

// Module ->
var Module = fx.Options(
	fx.Provide(NewUserService),
)
