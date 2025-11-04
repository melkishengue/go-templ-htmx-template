package server

import (
	"log/slog"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/melkishengue/gotemplate/cmd/frontend/views/layouts"
)

const (
	HXRequestHeader               = "HX-Request"
	HXHistoryRestoreRequestHeader = "HX-History-Restore-Request"
)

func requestsFullPage(c *gin.Context) bool {
	htmxRequest := c.GetHeader(HXRequestHeader) == "true"
	if !htmxRequest {
		return true
	}
	restoreRequest := c.GetHeader(HXHistoryRestoreRequestHeader) == "true"
	if restoreRequest {
		slog.Debug("restoring history")
	}
	return restoreRequest
}

func WithBase(c *gin.Context, component templ.Component, title, description string) templ.Component {
	return layouts.WithBase(component, title, description, requestsFullPage(c))
}
