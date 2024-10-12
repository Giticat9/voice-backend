package api

import "github.com/gin-gonic/gin"

type AppRouter struct {
	*gin.Engine
}

func CreateAppRouter() *AppRouter {
	return &AppRouter{gin.Default()}
}

func (router *AppRouter) InitAppRoutes() {
	router.InitAuthRoutes()
}
