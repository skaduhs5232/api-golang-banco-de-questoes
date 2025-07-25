package services

import (
	"encoding/json"
	"io"
	"net/http"
	"questoes/config"
	"questoes/structs"
	"strings"
)

func FiltrarQuestoes(w http.ResponseWriter, r *http.Request) {
	var questoes []structs.QuestaoSearch
	var filtros map[string]interface{}

	// Se for POST, lê o body com os filtros
	if r.Method == "POST" && r.ContentLength > 0 {
		err := json.NewDecoder(r.Body).Decode(&filtros)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
	}

	// Prepara os parâmetros para a função buscar_questoes
	params := map[string]interface{}{}

	if filtros != nil {
		// Filtra por enunciado
		if enunciado, ok := filtros["enunciado"].(string); ok && enunciado != "" {
			params["p_enunciado"] = enunciado
		}

		// Filtra por dificuldade
		if dificuldade, ok := filtros["dificuldade"].(string); ok && dificuldade != "" {
			params["p_dificuldade"] = dificuldade
		}

		// Filtra por assuntos
		if assuntosInterface, ok := filtros["assuntos"].([]interface{}); ok && len(assuntosInterface) > 0 {
			// Converte []interface{} para []string
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

	// Chama a função buscar_questoes do Supabase
	result := config.SupabaseClient.Rpc("buscar_questoes", "", params)

	err := json.Unmarshal([]byte(result), &questoes)

	if err != nil {
		http.Error(w, "Erro ao buscar questões: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questoes)
}

func AddQuestao(w http.ResponseWriter, r *http.Request) {
	// Lê o body bruto para debug
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
	}

	result := config.SupabaseClient.Rpc("inserir_questao", "", params)

	// Retorna o resultado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id":      result,
		"message": "Questão adicionada com sucesso",
	})
}
