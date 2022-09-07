package middleware

import (
	"practice-project/repository/jsonRepo"
	"practice-project/utils/jwt"
	"practice-project/utils/session"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	JsonLab jsonRepo.CashStorer
)

func JwtTokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var(
			jwtToken string
		)
		log.Println("")
		log.Println("[middileware] JwtTokenAuthentication is going to start !")
		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" {
			log.Printf("[middleware] empty Authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization header is required")
			return
		}
		tokenSlice := strings.Split(bearerToken, "Bearer ")
		if len(tokenSlice) == 1 {
			jwtToken = tokenSlice[0]
		}else{
			jwtToken = tokenSlice[1]
		}
		valid, err := jwt.ValidateToken(jwtToken, os.Getenv("jwt_secure_key"))
		if err != nil {
			log.Printf("[middleware] Token validation error : %s\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		if !valid {
			log.Printf("[middleware] Token is Invalide ")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token!")
			return
		}
		UserId, err := jwt.ValidateAndExtractUserID(jwtToken, os.Getenv("jwt_secure_key"))
		if err != nil {
			log.Printf("[middleware] Extracting UserId from Token UserId %s\n:", UserId)
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Somthing wents wrong!")
			return
		}
		if !JsonLab.IsObjectExist(UserId) {
			log.Printf("[middleware] got False when Checking object(tokenMetadata) in redis")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "session is terminated !, Can't Process a black listed token")
			return
		}
		log.Println("[middleware] JWT Token Authentication Successfully done :)")
		c.Next()
	}

}
func SessionAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("")
		log.Println("")
		log.Println("[middileware] SessionAuthentication is going to start")
		cookie, err := c.Request.Cookie("session-id")
		if err != nil {
			log.Printf("[middleware] Cookie extracting  %s\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, "You are Logged Out, please login again !")
			return
		}
		sessionToken := cookie.Value
		if sessionToken == "" {
			log.Printf("[middleware] sessionToken is Empty ")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "sessionToken is required. ")
			return
		}
		userSession, ok := session.IsTokenValid(sessionToken)
		if !ok {
			log.Printf("[middleware] Invalid SessionToken ")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Session not matched (invalid session). You will get access after login")
			return
		}
		expiry := userSession.IsSessionExpired()
		if expiry {
			log.Printf("[middleware] Session is expired by %v", userSession.Expiry.Sub(time.Now()))
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Session expired. please login again !")
			return
		}
		freshSessionToken := userSession.CreateNewSession()
		http.SetCookie(c.Writer,
			&http.Cookie{
				Name:    "session-id",
				Value:   freshSessionToken,
				Expires: time.Now().Add(time.Minute * 10),
			})
		log.Println("[middleware] Sessioin-Token Authentication Successfully done :)")
		c.Next()
	}

}
func LoginChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err1 := c.Request.Cookie("session-id")
		if err1 == nil {
			log.Printf("[LoginChecker] Cookie extracting (already login with session) %s\n", err1)
			c.AbortWithStatusJSON(http.StatusAlreadyReported, "already login with session !")
			return
		}
		_, err2 := c.Request.Cookie("token")
		if err2 == nil {
			log.Printf("[LoginChecker] Cookie extracting (already login with jwt) %s\n", err2)
			c.AbortWithStatusJSON(http.StatusAlreadyReported, "already login with jwt !")
			return
		}

		c.Next()
	}

}
