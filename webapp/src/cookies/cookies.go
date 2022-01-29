package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configurar utliza as variaves de ambiente para a criar a SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar registra as informações de autenticação
func Salvar(w http.ResponseWriter, ID, toke string) error {
	dados := map[string]string{
		"id":    ID,
		"token": toke,
	}

	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil

}

// Ler retorna os valores armazenados no cookies
func Ler(r *http.Request) (map[string]string, error) {
	cookies, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}
	valores := make(map[string]string)
	if erro = s.Decode("dados", cookies.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil

}
