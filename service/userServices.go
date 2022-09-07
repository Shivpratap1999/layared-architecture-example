package service

import (
	"log"
	"practice-project/error"
	"practice-project/models"

	"practice-project/utils/bcrypter"
)

func (u *userService) Registration(user *models.User) *error.Error {
	var (
		userCount int64
	)
	bcryptedPassword, err := bcrypter.GeneratePasswordHash(user.Password)
	if err != nil {
		log.Printf("[user-registration] error : %s while bcrypting Password  \n", err)
		return error.NewError(500 , "internal server error")
	}
	user.Password = bcryptedPassword
	if userCount, err = u.storer.CountUser(user.Email); err != nil {
		log.Printf("[user-registration] Database operation  error : %s\n", err)
		return error.NewError(500 , "internal server error")
	}
	if userCount > 0 {
		return nil
	}

	if err := u.storer.CreateUser(user); err != nil {
		log.Printf("[user-registration] Database Operation (Creating User) error : %s\n", err)
		return error.NewError(500 , "internal server error")
	}
	return nil
}

func (u *userService) FindAllUsers() ([]models.User, *error.Error) {
	var (
		users []models.User
	)
	if err := u.storer.GetAllUsers(&users); err != nil {
		log.Printf("[user-getAll] Database operation  error : %s\n", err)
		return nil, error.NewError(500 , "internal server error")
	}
	return users, nil
}
func (u *userService) FindUsersById(id int) (*models.User, *error.Error) {
	user, err := u.storer.FindUser(uint(id))
	if err != nil {
		log.Printf("[users-FindUser] Database operation  error : %s\n", err)
		return user, error.NewError(500 , "internal server error")
	}
	return user, nil
}
