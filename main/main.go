package main

import (
	"fmt"
	"net/http"
	"questoes/config"
	_ "questoes/controllers" // Importa para registrar os endpoints
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		allowedOrigin := "https://skaduhs5232.github.io"

		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	config.InitSupabase()

	mux := http.DefaultServeMux

	fmt.Println("Servidor rodando na porta 8080...")

	http.ListenAndServe(":8080", corsMiddleware(mux))
}
