package handlers

import (
	"practice-project/models"
	"practice-project/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	serve service.ServiceINF
}
func NewUserHandler(s service.ServiceINF)*UserHandler{
	return &UserHandler{serve:s}
}

func (uh *UserHandler) RegistrationController(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("[user] Unmarsling registration request data : ")
		c.JSON(http.StatusUnprocessableEntity, "Unprocessable entity provided !")
		return
	}

	err := uh.serve.Registration(user)
	if err != nil {
		log.Printf("[user] Unmarsling data : ")
		c.JSON(http.StatusInternalServerError, "Somthing wrong happened!")
	}
	msg := "Hello " + user.Name + " you are successfully registered"
	c.JSON(http.StatusOK, msg)

}

func (uh *UserHandler) GetUsersController(c *gin.Context) {
	users,err := uh.serve.FindAllUsers()
	if err != nil{
		c.JSON(http.StatusInternalServerError, "Somthing Wents Wrong!")
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetUserController(c *gin.Context){
	id,ok := c.Params.Get("id")
	if !ok {
		log.Println("[user-controler] fetching user's id from params")
		c.JSON(http.StatusBadRequest,"bad request, user'id is required")
		return
	}
	if id == "" {
		log.Println("[user-controler] validating id")
		c.JSON(http.StatusBadRequest,"user's id can't be a empty string")
		return
	}
	intUserId,err := strconv.Atoi(id)
	if err != nil{
		log.Printf("[user-controler] convering id string to int %s\n",err)
		c.JSON(http.StatusBadRequest,"id can't be a string")
		return
	}
	user,err := uh.serve.FindUsersById(intUserId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, "Somthing Wents Wrong!")
		return
	}
	c.JSON(http.StatusOK, user)
}
