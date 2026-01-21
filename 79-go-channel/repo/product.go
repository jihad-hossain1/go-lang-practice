package repo

import (
	"database/sql"
	"ecom/domain"
	"ecom/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}



// constructor or constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
	insert into products(
	title,
	description,
	price
	) values(
	 $1, 
	 $2, 
	 $3
	)
	 returning id
	`
	row := r.db.QueryRow(query,
		p.Title,
		p.Description,
		p.Price,
	)

	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *productRepo) Get(productId int) (*domain.Product, error) {
	var prd domain.Product

	query := `select * from products where id = $1`

	err := r.db.Get(&prd, query, productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return &prd, nil

}
func (r *productRepo) List(page,limit int64) ([]*domain.Product, error) {

	offset := ((page - 1) * limit) + 1

	var products []*domain.Product

	query := `select id, title, description, price from products limit $1 offset $2`

	err := r.db.Select(&products, query, limit, offset)

	fmt.Println(err)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return products, nil
}

func (r *productRepo) Count() (int64, error) {

	query := `select count(id) from products`

	var count int
	err := r.db.QueryRow(query).Scan(&count)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
	}

	return int64(count), nil
}

func (r *productRepo) Delete(productId int) error {
	query := `delete from products where id=$1`
	_, err := r.db.Exec(query, productId)
	if err != nil {
		return err
	}
	return nil
}
func (r *productRepo) Update(product domain.Product) (*domain.Product, error) {
	query := `
	update products
	set title = $1, description=$2, price=$3
	where id=$4
	`
	row := r.db.QueryRow(
		query,
		product.Title,
		product.Description,
		product.Price,
	)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
