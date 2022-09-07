package models

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Phone    int64  `json:"phone"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
