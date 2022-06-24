package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/karilho/API/controllers"
	"github.com/karilho/API/middleware"
)

// HandleRequest é quando eu receber o request na porta 8000 , nos endereços /, eu usar a o arquivo controolers para resposta, ou seja, demonstrar a func home
// que está localizada em controllers, logo, controllers.home.

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/API/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/API/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/API/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/API/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/API/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
