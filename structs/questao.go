package structs

type Questao struct {
	ID              string   `json:"id"`
	Enunciado       string   `json:"enunciado"`
	Alternativas    []string `json:"alternativas"`
	RespostaCorreta int      `json:"resposta_correta"`
	Dificuldade     string   `json:"dificuldade"`
	Assuntos        []string `json:"assuntos"`
	AnoQuestao      int      `json:"ano_questao"` // Novo campo adicionado
}

type QuestaoSearch struct {
	ID              string   `json:"ID"`
	Enunciado       string   `json:"Enunciado"`
	Alternativas    []string `json:"Alternativas"`
	RespostaCorreta int      `json:"RespostaCorreta"`
	Dificuldade     string   `json:"Dificuldade"`
	Assuntos        []string `json:"Assuntos"`
	AnoQuestao      int      `json:"Ano_questao"` // Novo campo adicionado
}
