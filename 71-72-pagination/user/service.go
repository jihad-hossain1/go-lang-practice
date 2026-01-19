package user

import (
	"ecom/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	return  svc.usrRepo.Create(user)

}
func (svc *service) Find(email string, pass string) (*domain.User, error) {
	return svc.usrRepo.Find(email, pass)
}
