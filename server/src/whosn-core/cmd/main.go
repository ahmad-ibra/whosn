package main

import (
	"context"
	"time"

	"github.com/Ahmad-Ibra/whosn-core/internal/config"
	"github.com/Ahmad-Ibra/whosn-core/internal/data"
	"github.com/Ahmad-Ibra/whosn-core/internal/endpoints"
	"github.com/Ahmad-Ibra/whosn-core/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	retry     = 10
	sleepTime = 1 * time.Second
)

func main() {
	// start db and run migrations
	pg, err := data.NewDB()
	if err != nil {
		panic(err)
	}

	// wait till we successfully ping the db or try to ping 10 times
	for i := 0; i < retry; i++ {
		err = pg.Conn.Ping(context.Background())
		if err == nil {
			break
		}
		time.Sleep(sleepTime)
	}
	if err != nil {
		panic(err)
	}

	router := initRouter(pg)
	router.Run(":" + config.GetConfig().Port)
}

func initRouter(db *data.PGStore) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())

	router.GET("/_hc", endpoints.Ping)
	apiV1 := router.Group("/api/v1")
	apiV1.Use(middleware.DB(db))
	{
		// register endpoints
		apiV1.POST("/user", endpoints.CreateUser)

		// login endpoint
		apiV1.POST("/login", endpoints.Login)

		apiV1Secured := apiV1.Group("/secured")
		apiV1Secured.Use(middleware.Auth())
		{
			// user endpoints
			apiV1Secured.DELETE("/user/:id", endpoints.DeleteUser)
			apiV1Secured.PUT("/user/:id", endpoints.UpdateUser)
			apiV1Secured.GET("/user/:id", endpoints.GetUser)

			// event endpoints
			apiV1Secured.GET("/events/owned", endpoints.ListOwnedEvents)   // TODO
			apiV1Secured.GET("/events/joined", endpoints.ListJoinedEvents) // TODO
			apiV1Secured.DELETE("event/:id", endpoints.DeleteEvent)        // TODO
			apiV1Secured.PUT("/event/:id", endpoints.UpdateEvent)          // TODO
			apiV1Secured.GET("/event/:id", endpoints.GetEvent)
			apiV1Secured.POST("/event", endpoints.CreateEvent)
			apiV1Secured.GET("/event/:id/join", endpoints.JoinEvent)   // TODO
			apiV1Secured.GET("/event/:id/leave", endpoints.LeaveEvent) // TODO
		}
	}
	return router
}
