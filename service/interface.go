package service

import (
	"practice-project/models"
	"practice-project/error"
)

type ServiceINF interface {
	Registration(user *models.User) *error.Error
	FindAllUsers() ([]models.User, *error.Error)
	FindUsersById(id int) (*models.User, *error.Error)
	LoginWithJwt(credential *models.LoginData) (string, *error.Error)
	LoginWithSession(credential *models.LoginData) (string, *error.Error)
	LogoutWithJwt(jwtToken string) *error.Error
	LogoutWithSession(sessionToken string)
}