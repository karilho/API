package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

// O var é um array personalidades, com uma struct Personalidade dentro
var Personalidades []Personalidade
