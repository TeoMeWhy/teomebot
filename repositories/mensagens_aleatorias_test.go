package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestLoadMessagensAleatorias(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	SetupMessageDB(db)

	repo := NewMensagensAleatoriasRepository()

	names := []string{
		"coach",
	}

	for _, v := range names {
		content, ok := repo.Mensagens[v]
		assert.Equal(t, true, ok)
		assert.NotEmpty(t, content.GetMensagem())
	}

}
