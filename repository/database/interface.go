package database
import "practice-project/models"

type Storer interface {
	CreateUser(u *models.User) error
	GetAllUsers(users *[]models.User) error
	FindUser(id uint) (*models.User ,error)
	CountUser(EmailId string) (int64, error)
	GetUserPass(EmailId string) (string, error)
}
