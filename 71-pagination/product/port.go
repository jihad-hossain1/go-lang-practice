package product

import ("ecom/domain"
prdctHandler "ecom/rest/handlers/product"
)

type Service interface {
	prdctHandler.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(productId int) error
	Update(p domain.Product) (*domain.Product, error)
}
