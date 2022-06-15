package main

import (
	"os"

	"github.com/Ahmad-Ibra/whosn-core/internal/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	router := initRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/_hc", endpoints.Ping)
	apiV1 := router.Group("/api/v1") //.Use(middlewares.Auth()).
	{
		// user endpoints
		apiV1.GET("/users", endpoints.ListUsers)
		apiV1.DELETE("/user/:id", endpoints.DeleteUser)
		apiV1.PUT("/user/:id", endpoints.UpdateUser)
		apiV1.GET("/user/:id", endpoints.GetUser)
		apiV1.POST("/user", endpoints.CreateUser)

		// event endpoints
		apiV1.GET("/events", endpoints.ListEvents)
		apiV1.DELETE("event/:id", endpoints.DeleteEvent)
		apiV1.PUT("/event/:id", endpoints.UpdateEvent)
		apiV1.GET("/event/:id", endpoints.GetEvent)
		apiV1.GET("/event/:id/join", endpoints.JoinEvent)
		apiV1.GET("/event/:id/leave", endpoints.LeaveEvent)
		apiV1.POST("/event", endpoints.CreateEvent)
	}
	return router
}
