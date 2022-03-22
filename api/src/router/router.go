package router

import (
	"api/src/rotas"

	"github.com/gorilla/mux"
)

// Gerar - vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
