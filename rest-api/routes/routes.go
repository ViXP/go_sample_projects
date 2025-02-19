package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	registerEventsRoutes(server)
	registerAuthsRoutes(server)
	registerUsersRoutes(server)
}

func registerEventsRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/", createEvent)
	authenticated.GET("/:id", getEvent)
	authenticated.PUT("/:id", updateEvent)
	authenticated.DELETE("/:id", deleteEvent)
}

func registerAuthsRoutes(server *gin.Engine) {
	server.POST("/register", signUpUser)
	server.POST("/login", signInUser)
}

func registerUsersRoutes(server *gin.Engine) {
}
