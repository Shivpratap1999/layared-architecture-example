package service

import (
	"log"
	"net/http"
	"os"
	"practice-project/error"
	"practice-project/models"
	"practice-project/utils/bcrypter"
	"practice-project/utils/jwt"
	"practice-project/utils/session"
	"time"
)

func (u *userService) LoginWithSession(credential *models.LoginData) (string, *error.Error) {
	PasswordHash, err := u.storer.GetUserPass(credential.Email)
	if err != nil {
		log.Printf("[login-] Database operation  %s\n", err)
		return "", error.NewError(500 , "internal server error")
	} else if PasswordHash == "" {
		log.Printf("[login] Email not exist in database")
		return "", error.NewError(http.StatusUnauthorized , "invalid credentials")
	}
	valid := bcrypter.ComparePasswordAndHash(credential.Password, PasswordHash)
	if !valid {
		log.Printf("[login] Password not matched (Unauthorised status)")
		return "", error.NewError(http.StatusUnauthorized , "invalid credentials")
	}
	session := session.NewClient(credential.Email)
	sessionToken := session.CreateNewSession()
	return sessionToken, nil
}


func (u *userService) LogoutWithSession(sessionToken string) {
	session.DestroySession(sessionToken)
}

// JWT_LOGIN_AND_LOGOUT_SERVICES>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
func (u *userService) LoginWithJwt(credential *models.LoginData) (string, *error.Error) {
	PasswordHash, err := u.storer.GetUserPass(credential.Email)
	if err != nil {
		log.Printf("[login-jwt] Database operation error: %s\n", err)
		return "", error.NewError(http.StatusInternalServerError , "somthing went wrong")
	} else if PasswordHash == "" {
		log.Printf("[login-jwt] Email not exist in database")
		return "", error.NewError(http.StatusUnauthorized , "invalid credentials")
	}
	valid := bcrypter.ComparePasswordAndHash(credential.Password, PasswordHash)
	if !valid {
		log.Printf("[login-jwt] Password not matched (Unauthorised status)")
		return "", error.NewError(http.StatusUnauthorized , "invalid credentials")

	}
	token, err := jwt.TokenProvider(credential.Email, os.Getenv("jwt_secure_key"))
	if err != nil {
		log.Printf("[login-jwt] error: %s while extracting token Token: %s\n", err, token)
		return "", error.NewError(http.StatusInternalServerError , "somthing went wrong")
	}
	u.cashStorer.StoreNewObject(credential.Email, token, 5*time.Minute)
	return token, nil
}

func (u *userService) LogoutWithJwt(jwtToken string) *error.Error {
	userinfo, err := jwt.ValidateAndExtractUserID(jwtToken, os.Getenv("jwt_secure_key"))
	if err != nil {
		log.Printf("[logout-jwt] error %s invalid tocken \n", err)
		return  error.NewError(http.StatusUnauthorized , "invalid token")
	}
	if err := u.cashStorer.DeleteObject(userinfo); err != nil {
		log.Printf("[login] error %s while deleting auth from redis\n", err)
		return error.NewError(http.StatusInternalServerError , "sonthing went wrong")
	}
	return nil

}
