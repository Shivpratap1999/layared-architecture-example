package routers

import (
	"practice-project/http/handlers"

	"github.com/gin-gonic/gin"
)
const(
	BASE_PATH string = "api/practice-project"
)
var (
	router *gin.Engine
	handler      handlers.UserHandlerInf
)

func InitialiseRouter(h handlers.UserHandlerInf) *gin.Engine {
	handler = h
	r := gin.Default()
	router = r
	StartUserRoutes()
	StartAuthRoutes()
	return r
}
