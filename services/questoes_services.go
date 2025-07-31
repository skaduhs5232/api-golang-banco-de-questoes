package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"questoes/config"
	"questoes/structs"
	"strings"
)

// FiltrarQuestoes busca questões com base nos filtros enviados no corpo da requisição.
func FiltrarQuestoes(w http.ResponseWriter, r *http.Request) {
	var questoes []structs.QuestaoSearch
	var filtros map[string]interface{}

	// Decodifica os filtros do corpo da requisição POST
	if r.Method == "POST" && r.ContentLength > 0 {
		err := json.NewDecoder(r.Body).Decode(&filtros)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
	}

	// Prepara os parâmetros para a chamada da função RPC "buscar_questoes"
	params := map[string]interface{}{}

	if filtros != nil {
		// Filtro por enunciado
		if enunciado, ok := filtros["enunciado"].(string); ok && enunciado != "" {
			params["p_enunciado"] = enunciado
		}

		// Filtro por dificuldade
		if dificuldade, ok := filtros["dificuldade"].(string); ok && dificuldade != "" {
			params["p_dificuldade"] = dificuldade
		}

		// Filtro por ano da questão
		// JSON decodifica números como float64, então precisamos fazer a conversão.
		if ano, ok := filtros["ano_questao"].(float64); ok && ano > 0 {
			params["p_ano_questao"] = int(ano) // Convertido para int
		}

		// Filtro por assuntos
		if assuntosInterface, ok := filtros["assuntos"].([]interface{}); ok && len(assuntosInterface) > 0 {
			assuntosStr := make([]string, 0, len(assuntosInterface))
			for _, v := range assuntosInterface {
				if str, ok := v.(string); ok && str != "" {
					assuntosStr = append(assuntosStr, str)
				}
			}
			if len(assuntosStr) > 0 {
				params["p_assuntos"] = assuntosStr
			}
		}
	}

	result := config.SupabaseClient.Rpc("buscar_questoes", "", params)

	err := json.Unmarshal([]byte(result), &questoes)
	if err != nil {
		http.Error(w, "Erro ao buscar questões: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questoes)
}

type SupabaseError struct {
	Code    string `json:"code"`
	Details string `json:"details"`
	Hint    string `json:"hint"`
	Message string `json:"message"`
}

func AddQuestao(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler body", http.StatusBadRequest)
		return
	}
	r.Body = io.NopCloser(strings.NewReader(string(body)))

	var novaQuestao structs.Questao
	err = json.NewDecoder(r.Body).Decode(&novaQuestao)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	params := map[string]interface{}{
		"p_enunciado":       novaQuestao.Enunciado,
		"p_alternativas":    novaQuestao.Alternativas,
		"p_respostacorreta": novaQuestao.RespostaCorreta,
		"p_dificuldade":     novaQuestao.Dificuldade,
		"p_assuntos":        novaQuestao.Assuntos,
		"p_ano_questao":     novaQuestao.AnoQuestao,
	}

	// Chama a função RPC do Supabase
	result := config.SupabaseClient.Rpc("inserir_questao", "", params)

	var supabaseErr SupabaseError
	if json.Unmarshal([]byte(result), &supabaseErr) == nil && supabaseErr.Code != "" {

		errorMsg := fmt.Sprintf("Erro do Supabase: %s (Hint: %s)", supabaseErr.Message, supabaseErr.Hint)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id":      result,
		"message": "Questão adicionada com sucesso",
	})
}
