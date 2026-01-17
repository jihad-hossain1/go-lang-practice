package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

// constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productId {
			return product, nil
		}
	}

	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productId int) error {
	var temptList []*Product

	for _, p := range r.productList {
		if p.ID != productId {
			temptList = append(temptList, p)
		}
	}

	r.productList = temptList

	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}

	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Mange",
		Description: "I Love Mango",
		Price:       110,
		ImgUrl:      "https://c8.alamy.com/comp/2N87A9W/mongo-fruit-hanging-on-tree-mongo-fruits-2N87A9W.jpg",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Mango juice",
		Description: "Fresh Mango Juice",
		Price:       100,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsx6-7fkbFv4qP3a9rsGy3zJSQIbeDamE2Pg&s",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Fresh organic source",
		Price:       10,
		ImgUrl:      "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/organic%20bananas.png?itok=_JpbRjWp-xPBdBLll",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)

}
