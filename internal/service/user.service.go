package service

import (
	"github.com/TS0906/go-ecommerce-backend-api/internal/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// us user service
func (us *UserService) GetInfoUser() string {
	return us.userRepo.GetInfoUser()
}
