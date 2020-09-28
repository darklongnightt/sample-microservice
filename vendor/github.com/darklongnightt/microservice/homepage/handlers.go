package homepage

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Handlers with logger as injected dependency
type Handlers struct {
	logger *log.Logger
}

// Home handler function
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	message := "Hello world"
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// Profile handler function
func (h *Handlers) Profile(w http.ResponseWriter, r *http.Request) {
	profile := &Profile{"Xavier", []string{"Calisthenics", "Coding"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// UploadFile handles file uploaded
func (h *Handlers) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("somefile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Header: ", header)
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(record)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Logger middleware that calculates processed time
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		h.logger.Printf("request processed in %vs\n", time.Now().Sub(startTime))
	}
}

// SetupRoutes routes handler functions to path related to homepage
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
	mux.HandleFunc("/profile", h.Logger(h.Profile))
	mux.HandleFunc("/upload", h.Logger(h.UploadFile))
}

// NewHandlers defines constructor for homepage handler
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{logger: logger}
}
