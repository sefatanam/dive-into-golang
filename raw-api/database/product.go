// Package database handle product list
package database

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var ProductList []Product

func init() {
	product1 := Product{
		ID:          1,
		Name:        "Product 1",
		Description: "Number 1 product",
	}

	product2 := Product{
		ID:          2,
		Name:        "Product 2",
		Description: "Number 2 product",
	}

	product3 := Product{
		ID:          3,
		Name:        "Product 3",
		Description: "Number 3 product",
	}
	ProductList = append(ProductList, product1)
	ProductList = append(ProductList, product2)
	ProductList = append(ProductList, product3)
}
