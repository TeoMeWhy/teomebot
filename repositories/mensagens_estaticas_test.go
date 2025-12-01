package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupMessageDB(db *gorm.DB) {

	db.AutoMigrate(&Messagem{})

	testMessages := []Messagem{
		{Chave: "cursos", Conteudo: "Plataforma de cursos livres: https://cursos.teomewhy.org"},
		{Chave: "livros", Conteudo: "Plataforma de livros: https://livros.teomewhy.org"},
	}

	db.Create(&testMessages)

}

func TestLoadMessagensEstaticas(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	SetupMessageDB(db)

	repo := NewMessageRepository(db)

	chaves := map[string]string{
		"cursos": "Plataforma de cursos livres: https://cursos.teomewhy.org",
		"livros": "Plataforma de livros: https://livros.teomewhy.org",
	}

	for k, v := range chaves {
		content, ok := repo.Messagens[k]
		assert.Equal(t, true, ok)
		assert.Equal(t, v, content.Conteudo)
	}

}
