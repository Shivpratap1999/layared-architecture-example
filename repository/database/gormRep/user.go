package database

import (
	"practice-project/models"
	"gorm.io/gorm"
)

type gormStorer struct {
	DB *gorm.DB
}

func NewGormStorer(db *gorm.DB) *gormStorer {
	return &gormStorer{DB: db}
}

func (gs gormStorer) CreateUser(u *models.User) error {
	err := gs.DB.Debug().Create(&u).Error
	return err
}
func (gs gormStorer) GetAllUsers(users *[]models.User) error {
	if err := gs.DB.Debug().Find(&users).Error; err != nil {
		return err
	}
	return nil
}
func (gs gormStorer) FindUser(id uint) (user *models.User, err error) {
	if err = gs.DB.Debug().First(&user, id).Error; err != nil {
		return nil, err
	}
	return
}
func (gs gormStorer) CountUser(EmailId string) (int64, error) {
	var userCount int64
	if err := gs.DB.Debug().Table("users").Where("email = ?", EmailId).Count(&userCount).Error; err != nil {
		return 0, err
	}
	return userCount, nil
}

func (us gormStorer) GetUserPass(EmailId string) (string, error) {
	var PasswordHash string
	if err := us.DB.Debug().Table("users").Where("email = ?", EmailId).Select("password").Order("id desc").Pluck("password", &PasswordHash).Error; err != nil {
		return "", err
	}
	return PasswordHash, nil
}
