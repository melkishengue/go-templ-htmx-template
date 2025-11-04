package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melkishengue/gotemplate/cmd/frontend/server"
	"github.com/melkishengue/gotemplate/cmd/frontend/views/contact"
	"github.com/melkishengue/gotemplate/cmd/frontend/views/home"
)

type Static struct {
}

func NewStatic() *Static {
	return &Static{}
}

func (*Static) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "", server.WithBase(c, home.Home(), "Home", "homepage"))
}

func (*Static) Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "", server.WithBase(c, contact.Contact(), "Contact", "contact"))
}

func (s *Static) Register(r *gin.RouterGroup) {
	r.Static("/public", "public")
	r.Static("/spec", "spec")

	r.GET("/", s.Home)
	r.GET("/contact", s.Contact)
	r.GET("/docs", func(c *gin.Context) {
		c.File("spec/documentation.html")
	})
}
