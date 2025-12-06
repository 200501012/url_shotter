package main

import (
	"log"
	"url_shotter/internal/config"
	"url_shotter/internal/http"
	"url_shotter/internal/http/domain"

	"github.com/lpernett/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	cfg := config.LoadConfig()

	db, err := cfg.ConnectDB()

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer cfg.CloseDB()

	// Executa migrações automáticas
	if err := db.AutoMigrate(&domain.URL{}); err != nil {
		log.Fatalf("Erro ao executar migrações: %v", err)
	}

	log.Println("✅ Conexão com o banco de dados estabelecida com sucesso!")
	log.Println("✅ Migrações executadas com sucesso!")

	app := http.NewServer(db)

	log.Println("Servidor iniciando na porta :8080")
	log.Fatal(app.Listen(":8080"))

}
