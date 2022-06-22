package main

import (
	"fmt"

	"github.com/karilho/API/routes"
)

// Aqui é só a mensagem pro meu terminal mesmo, o handlerequest esta em routes.
func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()

}
