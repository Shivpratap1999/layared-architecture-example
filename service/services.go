package service

import (
	"practice-project/repository/database"
	"practice-project/repository/jsonRepo"
)
type userService struct {
	storer     database.Storer
	cashStorer jsonRepo.CashStorer
}

func NewUserService(repo database.Storer, cs jsonRepo.CashStorer) *userService {
	return &userService{storer: repo, cashStorer: cs}
}
