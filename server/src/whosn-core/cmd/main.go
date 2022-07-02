package main

import (
	"context"

	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/Ahmad-Ibra/whosn-core/internal/endpoints"
	"github.com/Ahmad-Ibra/whosn-core/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// start db and run migrations
	db, err := data.NewDB()
	if err != nil {
		panic(err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	router := initRouter()
	router.Run(":" + config.GetConfig().Port)
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
			apiV1Secured.GET("/users", endpoints.ListUsers) // this is a temp endpoint for testing
			apiV1Secured.DELETE("/user/:id", endpoints.DeleteUser)
			apiV1Secured.PUT("/user/:id", endpoints.UpdateUser)
			apiV1Secured.GET("/user/:id", endpoints.GetUser)

			// event endpoints
			apiV1Secured.GET("/events", endpoints.ListEvents) // this is a temp endpoint for testing
			apiV1Secured.GET("/events/owned", endpoints.ListOwnedEvents)
			apiV1Secured.GET("/events/joined", endpoints.ListJoinedEvents)
			apiV1Secured.DELETE("event/:id", endpoints.DeleteEvent)
			apiV1Secured.PUT("/event/:id", endpoints.UpdateEvent)
			apiV1Secured.GET("/event/:id", endpoints.GetEvent)
			apiV1Secured.POST("/event", endpoints.CreateEvent)

			// eventUser endpoints
			apiV1Secured.GET("/event_users", endpoints.ListEventUsers) // this is a temp endpoint for testing
			apiV1Secured.GET("/event/:id/join", endpoints.JoinEvent)
			apiV1Secured.GET("/event/:id/leave", endpoints.LeaveEvent)
		}
	}
	return router
}
