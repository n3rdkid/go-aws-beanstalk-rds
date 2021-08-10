package controller

import (
	"go.uber.org/fx"
)

// Module ->
var Module = fx.Options(
	fx.Provide(NewUserController),
)
