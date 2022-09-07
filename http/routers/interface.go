package routers

import "github.com/gin-gonic/gin"

type Router interface{
	POST(uri string,f gin.HandlerFunc)
	GET(uri string, f gin.HandlerFunc)
	PUT(uri string, f gin.HandlerFunc)
	DELETE(uri string, f gin.HandlerFunc)
	RunServer(port string)
	UseMiddileware(f gin.HandlerFunc) 
}

