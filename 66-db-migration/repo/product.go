package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
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

func (r *productRepo) Create(p Product) (*Product, error) {
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

func (r *productRepo) Get(productId int) (*Product, error) {
	var prd Product

	query := `select * from products where id = $1`

	err := r.db.Get(&prd, query, productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return &prd, nil

}
func (r *productRepo) List() ([]*Product, error) {
	var products []*Product

	query := `select * form products`

	err := r.db.Select(&products, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}

	return products, nil
}

func (r *productRepo) Delete(productId int) error {
	query := `delete from products where id=$1`
	_, err := r.db.Exec(query, productId)
	if err != nil {
		return err
	}
	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
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
