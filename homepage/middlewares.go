package homepage

import (
	"net/http"
	"time"
)

// Logger middleware that calculates processed time
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		h.logger.Printf("request processed in %vs\n", time.Now().Sub(startTime))
	}
}
