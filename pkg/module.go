package pkg

import (
	"github.com/rickferrdev/sketch-service-email/pkg/mail"
	"go.uber.org/fx"
)

var Module = fx.Module("pkg", fx.Provide(
	mail.New,
))
