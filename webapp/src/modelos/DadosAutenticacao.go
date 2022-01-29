package modelos

// DadosAutenticacao contem os dados (ID, Token) do usuario autenticado
type DadosAutenticacao struct {
	ID    string `json: "id"`
	Token string `json: "token"`
}
