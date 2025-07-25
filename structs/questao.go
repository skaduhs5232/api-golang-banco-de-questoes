package structs

type Questao struct {
	ID              string   `json:"id"`
	Enunciado       string   `json:"enunciado"`
	Alternativas    []string `json:"alternativas"`
	RespostaCorreta int      `json:"resposta_correta"` // Certifique-se que está assim
	Dificuldade     string   `json:"dificuldade"`
	Assuntos        []string `json:"assuntos"`
}

type QuestaoSearch struct {
	ID              string   `json:"ID"`        // Note a diferença no mapeamento
	Enunciado       string   `json:"Enunciado"` // Para busca, usa os nomes do banco
	Alternativas    []string `json:"Alternativas"`
	RespostaCorreta int      `json:"RespostaCorreta"`
	Dificuldade     string   `json:"Dificuldade"`
	Assuntos        []string `json:"Assuntos"`
}
