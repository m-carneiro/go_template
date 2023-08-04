package health

import (
	"github.com/labstack/echo/v4"
	"github.com/programacao-fortificada/upsec/app"
	"net/http"

	"github.com/programacao-fortificada/upsec/app/health"
)

type handler struct {
	apps *app.Container
}

func Routes(group *echo.Group, apps *app.Container) {
	h := &handler{apps: apps}
	group.GET("/ping", h.Ping)
	group.GET("/health", echo.WrapHandler(h.Health()))
}

func (h *handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, health.Response{
		Status: "OK",
	})
}

func (h *handler) Health() http.Handler {
	return h.apps.Health.Health()
}
