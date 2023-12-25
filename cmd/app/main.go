package main

import (
	"encoding/json"
	"fmt"
	"irpf-ws/internal/database"
	"irpf-ws/internal/models"
	"log"
	"net/http"
	"os"
)

type ErrorMessage struct { 
	Message string `json:"message"`
	Status int `json:"status"`
	ErrorMessage string `json:"error"` 
}

type Message struct {
	Text string `json:"message"`
}

func main() {
	_, err := database.InitDb()
	if err != nil {
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusInternalServerError)
			var message string = "Método não permitido"
			var status int = http.StatusMethodNotAllowed
			var errorMessage string = "Method not allowed"
			json.NewEncoder(w).Encode(ErrorMessage{Message: message, Status: status, ErrorMessage: errorMessage})
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "Seja bem-vindo(a) à API IRPF.")
	})

	http.HandleFunc("/contribuintes/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			var message string = "Método não permitido."
			var status int = http.StatusMethodNotAllowed
			var errorMessage string = "Method not allowed"
			json.NewEncoder(w).Encode(ErrorMessage {Message: message, Status: status, ErrorMessage: errorMessage})
			return
		}

		contribuintes, err := database.GetContribuintes()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			var message string = "Erro ao buscar dados dos Contribuintes"
			var status int = 500
			var errorMessage string = err.Error()
			json.NewEncoder(w).Encode(ErrorMessage{Message: message, Status: status, ErrorMessage: errorMessage})
			return
		}

		if contribuintes == nil {
			contribuintes = make([]models.Contribuinte, 0)
		}
		
		json.NewEncoder(w).Encode(contribuintes)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server started in port %s", port)

	http.ListenAndServe(":" + port, nil)
}