package handlers

import "github.com/gin-gonic/gin"

type UserHandlerInf interface {
	RegistrationController(c *gin.Context)
	GetUsersController(c *gin.Context)
	GetUserController(c *gin.Context)
	LoginJwtController(c *gin.Context)
	LogoutJwtController(c *gin.Context)
	LoginSessionController(c *gin.Context)
	LogoutSessionController(c *gin.Context)
}
