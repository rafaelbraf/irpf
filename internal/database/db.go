package database

import (
	"database/sql"
	"fmt"
	"irpf-ws/internal/models"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() (*sql.DB, error) {
	conn, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db = conn
	fmt.Printf("Conexão com o Banco de Dados realizada com sucesso!")
	
	return db, nil
}

func GetDb() *sql.DB {
	return db
}

func GetContribuintes() ([]models.Contribuinte, error) {
	rows, err := db.Query("SELECT id, cpf, nome, celular, data_nascimento, email, endereco, natureza_ocupacao, ocupacao_principal FROM contribuintes")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var contribuintes []models.Contribuinte

	for rows.Next() {
		var contribuinte models.Contribuinte

		err:= rows.Scan(
			&contribuinte.ID, 
			&contribuinte.Cpf, 
			&contribuinte.Nome, 
			&contribuinte.Celular, 
			&contribuinte.DataNascimento, 
			&contribuinte.Email, 
			&contribuinte.Endereco,
			&contribuinte.NaturezaOcupacao,
			&contribuinte.OcupacaoPrincipal)
		if err != nil {
			return nil, err
		}

		contribuintes = append(contribuintes, contribuinte)
	}

	return contribuintes, nil
}

func getDBConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}