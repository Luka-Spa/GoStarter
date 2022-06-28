package controller

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type httpRouter struct {
	router *gin.Engine
	api    *gin.RouterGroup
}

func NewHTTPRouter() *httpRouter {
	router := gin.Default()
	router.SetTrustedProxies([]string{})
	return &httpRouter{
		router: router,
		api:    router.Group("/api"),
	}
}

func (r *httpRouter) Run() {
	port, hasValue := os.LookupEnv("HTTP_PORT")
	if !hasValue {
		port = "8080"
	}
	r.router.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), port))
}
