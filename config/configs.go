package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

func InitSupabase() {
	// Tenta carregar o arquivo .env, mas não trata a ausência dele como um erro fatal.
	// Em produção (como na Render), as variáveis virão do ambiente.
	godotenv.Load()

	// Pega as variáveis de ambiente
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("SUPABASE_URL e SUPABASE_KEY devem estar definidas no ambiente")
	}

	// Inicializa o cliente Supabase
	var err error
	SupabaseClient, err = supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatal("Erro ao conectar com Supabase:", err)
	}

	log.Println("Conectado ao Supabase com sucesso!")
}
