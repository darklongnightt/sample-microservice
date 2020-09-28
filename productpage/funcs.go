package productpage

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/darklongnightt/microservice/db"
)

func (h *Handlers) createProduct(w http.ResponseWriter, r *http.Request) {
	// Try to decode the request body into product struct
	product := &db.Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert into db table: product_items
	if err := h.db.Insert(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode result into []bytes(json)
	js, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func (h *Handlers) getProduct(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, `missing query param: "id"`, http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(keys[0], 10, 32)
	if err != nil {
		http.Error(w, `"id" must be of type int`, http.StatusBadRequest)
		return
	}

	product := &db.Product{ID: int(id)}
	if err := h.db.Model(product).WherePK().Select(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
