package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositoDePublicacoes cria um repositorio de publicações
func NovoRepositoDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar recebe uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES(?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}
