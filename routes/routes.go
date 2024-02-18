package routes

import (
	"log"
	"net/http"

	"github.com/bb9leko/apiRest-go-alura/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	/*http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(`:8080`, nil))*/

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(`:8080`, r))

}
