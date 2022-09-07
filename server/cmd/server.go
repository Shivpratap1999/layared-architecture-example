package main

import (
	"practice-project/server/settings"
	"flag"
	"os"

	"github.com/gin-contrib/cors"
)

func main() {
	os.Setenv("jwt_secure_key", "super-secure-key")
	port := flag.String("port", ":8080", "port-number")
	flag.Parse()

	r := settings.ConfigurServerSettings()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	r.Use(cors.New(corsConfig))
	r.Run(*port)
}
