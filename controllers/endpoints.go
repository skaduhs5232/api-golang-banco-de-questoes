package controllers

import (
	"net/http"
	"questoes/services"
)

func init() {
	http.HandleFunc("POST /questoes", services.FiltrarQuestoes)

	http.HandleFunc("POST /add-questao", services.AddQuestao)
}
