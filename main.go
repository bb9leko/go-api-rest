package main

import (
	"fmt"

	"github.com/bb9leko/apiRest-go-alura/models"
	"github.com/bb9leko/apiRest-go-alura/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{

		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Hello World!")
	routes.HandleRequest()

}
