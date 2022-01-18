package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger escreve informaçĩoes da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado.
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando....")
		next(w, r)
	}

}
