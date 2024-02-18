package models

type Personalidade struct {

	// `json:"nome"` - Como vamos deserializar
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

// mockando os dados (ao invés de ir no banco...)
var Personalidades []Personalidade
