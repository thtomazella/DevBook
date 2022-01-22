package controllers

import "net/http"

// CriarPublicacao adiciona uma nova publicacao no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {}

// BuscarPublicacoes  traz as publicacões que apareceriam no feed do usuarios
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {}

// BuscarPublicacao traz uma única publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {}

// AtualizarPublicacao altera os dados de uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {}

// DeletarPublicacao exclui os dados de uma publicação
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {}
