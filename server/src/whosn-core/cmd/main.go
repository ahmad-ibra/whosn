package main

import (
	"net/http"
	"os"

	"github.com/Ahmad-Ibra/whosn-core/internal/endpoints"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/events", endpoints.ListEvents)
	router.HandleFunc("/api/v1/event/{id}", endpoints.DeleteEvent).Methods("DELETE")
	router.HandleFunc("/api/v1/event/{id}", endpoints.UpdateEvent).Methods("PUT")
	router.HandleFunc("/api/v1/event/{id}", endpoints.GetEvent)
	router.HandleFunc("/api/v1/event/{id}/join", endpoints.JoinEvent)
	router.HandleFunc("/api/v1/event/{id}/leave", endpoints.LeaveEvent)
	router.HandleFunc("/api/v1/event", endpoints.CreateEvent).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Infof("Starting http server on port %v", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

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
	apiV1 := router.Group("/api/v1")
	{
		apiV1.DELETE("/user/:id", endpoints.DeleteUser)
		apiV1.PUT("/user/:id", endpoints.UpdateUser)
		apiV1.GET("/user/:id", endpoints.GetUser)
		apiV1.POST("/user", endpoints.CreateUser)
		//apiV1.POST("/token", controllers.GenerateToken)
		//apiV1.POST("/user/register", controllers.RegisterUser)
		//secured := apiV1.Group("/secured").Use(middlewares.Auth())
		//{
		//	secured.GET("/ping", controllers.Ping)
		//}
	}
	return router

}
