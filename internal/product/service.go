package product

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	oid := primitive.NewObjectID()
	product := Product{
		ID:        oid,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	json.NewDecoder(r.Body).Decode(&product)
	Create(product)
	json.NewEncoder(w).Encode(map[string]any{"result": "Producto guardado correctamente"})
}

func (s *Handler) ReadProducts(w http.ResponseWriter, r *http.Request) {
	products, err := Read()
	if err != nil {
		panic(err)
	}

	if len(products) == 0 {
		json.NewEncoder(w).Encode(map[string]any{"result": "No hay Productos"})
	} else {
		json.NewEncoder(w).Encode(products)
	}
}

func (s *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product := Product{}
	json.NewDecoder(r.Body).Decode(&product)
	Update(product, params["id"])
	json.NewEncoder(w).Encode(map[string]any{"result": "Producto actualizado correctamente"})
}

func (s *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Delete(params["id"])
	json.NewEncoder(w).Encode(map[string]any{"result": "Producto eliminado correctamente"})

}
