package modelos

// DadosAutenticacao contem o token e o ID do usuario autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"Token"`
}
