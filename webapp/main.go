package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {

	utils.CarregarTemplates()

	r := router.Gerar()

	fmt.Println("Escutando na Porta 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
