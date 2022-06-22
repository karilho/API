package controllers

import (
	"fmt"
	"net/http"
)

// Home é para quando eu receber um pedido, eu quero demonstrar a mensagem pagina inicial.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Página INICIAL")
}
