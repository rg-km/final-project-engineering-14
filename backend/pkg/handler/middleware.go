package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

func (handler *Handler) AllowOrigin(w http.ResponseWriter, req *http.Request) {
	// localhost:8080 origin mendapat ijin akses
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// allow cookie
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *Handler) POST(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.AllowOrigin(w, r)
		encoder := json.NewEncoder(w)
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			encoder.Encode(web.WebResponse{
				Status:  http.StatusMethodNotAllowed,
				Message: "Method not allowed",
				Data:    "Need POST method",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
