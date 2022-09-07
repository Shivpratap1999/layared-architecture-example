package routers

import (
	"log"
	"github.com/gin-gonic/gin"
)

type ginRouter struct {
	router *gin.Engine
}

func NewGinRouter() *ginRouter {
	r := gin.Default()
	return &ginRouter{router: r}
}

func (g *ginRouter) POST(uri string, f gin.HandlerFunc) {
	g.router.POST(uri, f)
}

func (g *ginRouter) GET(uri string, f gin.HandlerFunc) {
	g.router.GET(uri, f)
}

func (g *ginRouter) PUT(uri string, f gin.HandlerFunc) {
	g.router.PUT(uri, f)
}

func (g *ginRouter) DELETE(uri string, f gin.HandlerFunc) {
	g.router.DELETE(uri, f)
}

func (g *ginRouter) RunServer(port string) {
	log.Println("Gin Server is running on Port:", port)
	g.router.Run(port)
}
func (g *ginRouter)UseMiddileware(f gin.HandlerFunc){
	g.router.Use(f)
}