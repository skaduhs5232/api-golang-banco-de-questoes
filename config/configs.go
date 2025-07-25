package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

func InitSupabase() {
    // Carrega variáveis do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar arquivo .env")
    }

    // Pega as variáveis de ambiente
    supabaseURL := os.Getenv("SUPABASE_URL")
    supabaseKey := os.Getenv("SUPABASE_KEY")

    if supabaseURL == "" || supabaseKey == "" {
        log.Fatal("SUPABASE_URL e SUPABASE_KEY devem estar definidas no .env")
    }

    // Inicializa o cliente Supabase
    SupabaseClient, err = supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{})
    if err != nil {
        log.Fatal("Erro ao conectar com Supabase:", err)
    }

    log.Println("Conectado ao Supabase com sucesso!")
}