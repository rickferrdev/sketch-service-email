package config

import (
	"github.com/rickferrdev/sketch-service-email/config/env"
	"go.uber.org/fx"
)

var Module = fx.Module("config", fx.Provide(
	env.New,
))
