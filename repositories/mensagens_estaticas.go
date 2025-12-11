package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type Messagem struct {
	Chave    string `gorm:"primaryKey"`
	Conteudo string
}

type MessageRepository struct {
	ConDB     *gorm.DB
	Messagens map[string]Messagem
}

func (r *MessageRepository) LoadMessagensEstaticas() {

	var mensagens []Messagem
	r.ConDB.Find(&mensagens)

	r.Messagens = map[string]Messagem{}
	for _, msg := range mensagens {
		r.Messagens[msg.Chave] = msg
	}

}

func (r *MessageRepository) CreateMessagem(chave string, conteudo string) error {
	msg := &Messagem{Chave: chave, Conteudo: conteudo}
	if err := r.ConDB.Create(&msg).Error; err != nil {
		return err
	}
	r.Messagens[chave] = *msg
	return nil
}

func (r *MessageRepository) UpdateMessagem(chave string, conteudo string) error {
	msg, exists := r.Messagens[chave]
	if !exists {
		return fmt.Errorf("err: {%s} não é uma chave existente", chave)
	}
	msg.Conteudo = conteudo
	if err := r.ConDB.Save(&msg).Error; err != nil {
		return err
	}
	r.Messagens[chave] = msg
	return nil
}

func (r *MessageRepository) ShowMessage(key string) string {
	msg, exists := r.Messagens[key]
	if !exists {
		return ""
	}
	return msg.Conteudo
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	repo := &MessageRepository{ConDB: db}
	repo.LoadMessagensEstaticas()
	return repo
}
