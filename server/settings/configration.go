package settings

import (
	"log"
	"practice-project/http/handlers"
	"practice-project/http/middleware"
	database "practice-project/repository/database/gormRep"
	jsonRepo "practice-project/repository/jsonRepo/redis"
	"practice-project/http/routers"
	"practice-project/service"

	"github.com/gin-gonic/gin"
)

func ConfigurServerSettings() *gin.Engine {
	mysqlGormDB, err := ConnectGormDB()
	if err != nil {
		log.Fatal("Error when configuring Gorm DataBase")
	}
	redisClint, err := ConfigureRedis()
	if err != nil {
		log.Fatal("Error when configuring Redis Server")
	}

	storer := database.NewGormStorer(mysqlGormDB)

	chashStorer := jsonRepo.NewredisStorer(redisClint)

	middleware.JsonLab = chashStorer

	serviceInf := service.NewUserService(storer, chashStorer)

	userHandler := handlers.NewUserHandler(serviceInf)

	router := routers.InitialiseRouter(userHandler)

	return router
}
