package homepage

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handlers with logger as injected dependency
type Handlers struct {
	logger *log.Logger
}

// NewHandlers defines constructor for homepage handler
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{logger: logger}
}

// SetupRoutes routes handler functions to path related to homepage
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
	mux.HandleFunc("/profile", h.Logger(h.Profile))
	mux.HandleFunc("/upload", h.Logger(h.UploadFile))
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
	profile := h.getProfile()
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
	file, _, err := r.FormFile("somefile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.readFile(file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
