package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/karilho/API/models"
)

// Este arquivo gerenciará o que será MOSTRADO, adicionar todas funções mães aqui.
// Home é para quando eu receber um pedido, eu quero demonstrar a mensagem pagina inicial.

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Página INICIAL")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
