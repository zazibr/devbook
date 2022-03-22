package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// para criar
//  go mod init nesse caso api
//  go mod init api
//
// instalar o mux
// go get github.com/gorilla/mux
//
func main() {
	config.Carregar()

	fmt.Printf("Escutando na porta %d", config.Porta)

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
