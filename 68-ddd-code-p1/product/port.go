package product

import "ecom/domain"

type Service struct {
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(productId int) error
	Update(p domain.Product) (*domain.Product, error)
}
