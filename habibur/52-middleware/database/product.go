package database

/*
- when struct define key First character small letter then not accessble other package ( this are the private method to allow intier main package only)
- we define this way `json:"id"`
*/

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var ProductList []Product

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Mange",
		Description: "I Love Mango",
		Price:       110,
		ImgUrl:      "https://c8.alamy.com/comp/2N87A9W/mongo-fruit-hanging-on-tree-mongo-fruits-2N87A9W.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Mango juice",
		Description: "Fresh Mango Juice",
		Price:       100,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsx6-7fkbFv4qP3a9rsGy3zJSQIbeDamE2Pg&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Fresh organic source",
		Price:       10,
		ImgUrl:      "https://www.dole.com/sites/default/files/styles/1024w768h-80/public/media/2025-01/organic%20bananas.png?itok=_JpbRjWp-xPBdBLll",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)

}
