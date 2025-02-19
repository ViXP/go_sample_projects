package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	useEventsRoutes(server)
	useAuthRoutes(server)
	useRegistrationsRoutes(server)
}

func useAuthRoutes(server *gin.Engine) {
	server.POST("/signup", signUpUser)
	server.POST("/signin", signInUser)
}

func useEventsRoutes(server *gin.Engine) {
	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/", createEvent)
	authenticated.PUT("/:id", updateEvent)
	authenticated.DELETE("/:id", deleteEvent)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
}

func useRegistrationsRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events/:id/register", createRegistration)
	authenticated.DELETE("/events/:id/register", deleteRegistration)
}
