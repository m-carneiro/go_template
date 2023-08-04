package app

import (
	"github.com/programacao-fortificada/upsec/app/health"
	"github.com/programacao-fortificada/upsec/internal/config"
	"github.com/programacao-fortificada/upsec/pkg/log"
)

type Container struct {
	Health health.App
}

type Options struct {
	Cfg *config.Config
	Log *log.Log
}

func New(opts Options) *Container {
	return &Container{
		Health: health.New(opts.Cfg),
	}
}
