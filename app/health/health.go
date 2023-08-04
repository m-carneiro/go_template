package health

import (
	"github.com/hellofresh/health-go/v5"
	healthHttp "github.com/hellofresh/health-go/v5/checks/http"
	"github.com/programacao-fortificada/upsec/internal/config"
	"net/http"
	"time"
)

type App interface {
	Health() http.Handler
}

type app struct {
	Cfg *config.Config
}

func New(cfg *config.Config) App {
	return &app{
		Cfg: cfg,
	}
}

func (a *app) Health() http.Handler {
	h, _ := health.New(health.WithSystemInfo())

	_ = h.Register(health.Config{
		Name:      "database",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: healthHttp.New(healthHttp.Config{
			URL: "https://www.google.com",
		}),
	})

	return h.Handler()
}
