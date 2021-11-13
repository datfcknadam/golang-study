package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}
	v1 := r.router.Group("/v1")
	r.addName(v1)
	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}
