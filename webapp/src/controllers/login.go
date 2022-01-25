package controllers

import "net/http"

/// CarregarTelaDeLogin vai carregar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de login"))
}
