// Package handlers requests.
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sefatanam/rga/database"
	"github.com/sefatanam/rga/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)
}
