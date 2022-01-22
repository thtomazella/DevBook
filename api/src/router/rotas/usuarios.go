package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{ // Cadastro o usuario
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{ // Busca Todos os Usuários
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},

	{ // Buscar somente um usuario por ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},

	{ //Alterar o usuario por ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoUsuario,
		RequerAutenticacao: true,
	},
	{ // Deletar o registro por ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
}
