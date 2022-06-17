package main

import (
	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/Ahmad-Ibra/whosn-core/internal/endpoints"
	"github.com/Ahmad-Ibra/whosn-core/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()
	router := initRouter()
	router.Run(":" + cfg.Port)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())

	router.GET("/_hc", endpoints.Ping)
	apiV1 := router.Group("/api/v1")
	{
		// login endpoint
		apiV1.POST("/login", endpoints.Login)

		// register endpoints
		apiV1.POST("/user", endpoints.CreateUser)

		apiV1Secured := apiV1.Group("/secured").Use(middleware.Auth())
		{
			// user endpoints
			apiV1Secured.GET("/users", endpoints.ListUsers).Use(middleware.Auth())
			apiV1Secured.DELETE("/user/:id", endpoints.DeleteUser).Use(middleware.Auth())
			apiV1Secured.PUT("/user/:id", endpoints.UpdateUser).Use(middleware.Auth())
			apiV1Secured.GET("/user/:id", endpoints.GetUser).Use(middleware.Auth())

			// event endpoints
			apiV1Secured.GET("/events", endpoints.ListEvents).Use(middleware.Auth())
			apiV1Secured.DELETE("event/:id", endpoints.DeleteEvent).Use(middleware.Auth())
			apiV1Secured.PUT("/event/:id", endpoints.UpdateEvent).Use(middleware.Auth())
			apiV1Secured.GET("/event/:id", endpoints.GetEvent).Use(middleware.Auth())
			apiV1Secured.GET("/event/:id/join", endpoints.JoinEvent).Use(middleware.Auth())
			apiV1Secured.GET("/event/:id/leave", endpoints.LeaveEvent).Use(middleware.Auth())
			apiV1Secured.POST("/event", endpoints.CreateEvent).Use(middleware.Auth())
		}
	}
	return router
}
