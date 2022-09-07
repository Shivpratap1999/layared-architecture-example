package handlers

import (
	"practice-project/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) LoginSessionController(c *gin.Context) {
	var (
		credential models.LoginData
	)
	_, err := c.Request.Cookie("session-id")
	if err == nil {
		log.Printf("[login-controller] Cookie extracting (allready login)  %s\n", err)
		c.AbortWithStatusJSON(http.StatusAlreadyReported, "You are already logged in ?")
		return
	}
	if err := c.ShouldBindJSON(&credential); err != nil {
		log.Printf("[login-Controller] Unmarsling data : ")
		c.JSON(http.StatusUnprocessableEntity, "Unprocessable entity provided !")
		return
	}
	if credential.Email == "" || credential.Password == "" {
		c.JSON(http.StatusBadRequest, "Email or Password is empty!")
		return
	}
	token, serviceError := uh.serve.LoginWithSession(&credential)
	if serviceError != nil {
		log.Printf("[login-Controller] error :%s while logging with session \n: ", serviceError.Error())
		c.JSON(serviceError.Code(), serviceError.Error())
		return
	}
	http.SetCookie(c.Writer,
		&http.Cookie{
			Name:    "session-id",
			Value:   token,
			Expires: time.Now().Add(time.Minute * 10),
		})
	respMap := make(map[string]interface{})
	respMap["status"] = "Successfully Login by session"
	c.JSON(http.StatusOK, respMap)

}
func (uh *UserHandler) LogoutSessionController(c *gin.Context) {
	cookie, err := c.Request.Cookie("session-id")
	if err != nil {
		log.Println("[Logout-Controller] fetching token cookies error:", err)
		c.JSON(http.StatusBadRequest, "already logged out")
	}
	sessionToken := cookie.Value
	if sessionToken == "" {
		log.Printf("[logout-controller] sessionToken is Empty ")
		c.JSON(http.StatusUnauthorized, "sessionToken is required. ")
		return
	}
	uh.serve.LogoutWithSession(sessionToken)
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.MaxAge = -1
	http.SetCookie(c.Writer, cookie)
	respMap := make(map[string]interface{})
	respMap["status"] = "Successfully logout from session"
	c.JSON(http.StatusOK, respMap)
}

//JWT_TOKEN_LOGIN_LOGOUT>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func (uh *UserHandler) LoginJwtController(c *gin.Context) {
	var (
		credential models.LoginData
	)
	if err := c.ShouldBindJSON(&credential); err != nil {
		log.Printf("[login-Controller] Unmarsling data : ")
		c.JSON(http.StatusUnprocessableEntity, "Unprocessable entity provided !")
		return
	}
	if credential.Email == "" || credential.Password == "" {
		c.JSON(http.StatusBadRequest, "Email or Password is empty!")
		return
	}
	token, serviceError := uh.serve.LoginWithJwt(&credential)
	if serviceError != nil {
		log.Printf("[login-Controller] error: %s while logging with jwt error-code %d\n: ", serviceError.Error(),serviceError.Code())
		c.JSON(serviceError.Code(), serviceError.Error())
		return
	}
	respMap := make(map[string]interface{})
	respMap["status"] = "Successfully Login by jwt token"
	respMap["token"] = token
	c.JSON(http.StatusOK, respMap)

}

func (uh *UserHandler) LogoutJwtController(c *gin.Context) {
	var(
		jwtToken string
	)
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		c.JSON(http.StatusUnauthorized, "Authorization header is required")
		return
	}
	tokenSlice := strings.Split(bearerToken, "Bearer ")
	if len(tokenSlice) == 1 {
		jwtToken = tokenSlice[0]
	}else{
		jwtToken = tokenSlice[1]
	}
	serviceError := uh.serve.LogoutWithJwt(jwtToken)
	if serviceError != nil {
		log.Printf("[login-Controller] error: %s while logout with jwt error-code %d\n: ", serviceError.Error(),serviceError.Code())
		c.JSON(serviceError.Code(), serviceError.Error())
		return
	}
	respMap := make(map[string]interface{})
	respMap["status"] = "Successfully logout from jwt"
	c.JSON(http.StatusOK, respMap)
}
