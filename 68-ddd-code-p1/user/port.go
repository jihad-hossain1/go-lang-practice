package user

import (
	"ecom/domain"
	userHandler "ecom/rest/handlers/user"
)

type Service interface {
	userHandler.Service // embedding
}

type UserRepo interface {
	Create(usr domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
