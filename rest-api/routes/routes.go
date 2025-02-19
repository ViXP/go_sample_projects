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
	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/", createEvent)
	authenticated.PUT("/:id", updateEvent)
	authenticated.DELETE("/:id", deleteEvent)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
}

func registerAuthsRoutes(server *gin.Engine) {
	server.POST("/register", signUpUser)
	server.POST("/login", signInUser)
}

func registerUsersRoutes(server *gin.Engine) {
}
