package main

import (
	"fmt"

	"github.com/karilho/API/models"
	"github.com/karilho/API/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "P1", Historia: "H1"},
		{Id: 2, Nome: "P2", Historia: "H2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()

}
