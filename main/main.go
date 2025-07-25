package main

import (
    "fmt"
    "net/http"
    "questoes/config"
    _ "questoes/controllers"
)

func main() {
    // Inicializa a conexão com Supabase
    config.InitSupabase()
    
    fmt.Println("Servidor rodando na porta 8080...")
    http.ListenAndServe(":8080", nil)
}