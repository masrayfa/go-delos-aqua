package helper

import (
	"log"
	"net/http"
)

func StoreOriginalPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simpan path asli ke dalam context request atau variabel global
		log.Println("@Middleware.StoreOriginalPath:r.URL.Path -> ", r.URL.Path)
		r.Header.Set("X-Original-Path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}