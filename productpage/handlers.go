package productpage

import (
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg"
)

// Handlers ...
type Handlers struct {
	logger *log.Logger
	db     *pg.DB
}

// NewHandlers ...
func NewHandlers(logger *log.Logger, db *pg.DB) *Handlers {
	return &Handlers{
		logger,
		db,
	}
}

// SetupRoutes ...
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/product", h.Logger(h.Product))
}

// Logger middleware that calculates processed time
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
		h.logger.Printf("request processed in %vs\n", time.Now().Sub(startTime))
	}
}

// Product handler func creates new product in the db
func (h *Handlers) Product(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.createProduct(w, r)
	} else if r.Method == http.MethodGet {
		h.getProduct(w, r)
	}
}
