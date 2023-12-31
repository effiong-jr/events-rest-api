package routes

import "github.com/gin-gonic/gin"

func UsersRoutes(server *gin.Engine) {
	server.POST("/signup", signupUserHandler)
	server.POST("/login", loginHandler)
}
