package routes

import (
	"log"
	"net/http"

	"github.com/karilho/API/controllers"
)

// HandleRequest é quando eu receber o request na porta 8000 , nos endereços /, eu usar a função HOME para resposta, ou seja, demonstrar a func home
// que está localizada em controllers, logo, controllers.home.
func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
