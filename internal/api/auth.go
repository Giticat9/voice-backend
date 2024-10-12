package api

import handlers "voice/main/internal/handlers/auth"

func (router *AppRouter) InitAuthRoutes() {
	authGroup := router.Group("/api/auth")
	authGroup.POST("/login", handlers.LoginHandler)
}
