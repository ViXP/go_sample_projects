package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	registerEventsRoutes(server)
	registerAuthsRoutes(server)
	registerUsersRoutes(server)
}

func registerEventsRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
}

func registerAuthsRoutes(server *gin.Engine) {
	server.POST("/register", signUpUser)
	server.POST("/login", signInUser)
}

func registerUsersRoutes(server *gin.Engine) {
}
