package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL respresenta a URL para comunicação com API
	APIURL = ""
	//Porta onde a aplicação está rodando
	Porta = 0
	//HashKey é utilizado para autenticar o cookie
	HashKey []byte
	// BlockKey é utilizado para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORTA"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
