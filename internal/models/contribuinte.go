package models

import "database/sql"

type Contribuinte struct {
	ID                int     `json:"id"`
	Cpf               string  `json:"cpf"`
	Nome              string  `json:"nome"`
	Celular           string  `json:"celular"`
	Endereco          string  `json:"endereco"`
	DataNascimento    sql.NullString `json:"dataNascimento"`
	Email             string  `json:"email"`
	NaturezaOcupacao  string  `json:"naturezaOcupacao"`
	OcupacaoPrincipal string  `json:"ocupacaoPrincipal"`
}