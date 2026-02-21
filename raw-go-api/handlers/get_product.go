package handlers

import (
	"net/http"

	"github.com/sefatanam/rga/database"
	"github.com/sefatanam/rga/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200)
}
