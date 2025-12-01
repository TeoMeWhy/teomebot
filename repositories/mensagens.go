package repositories

import "gorm.io/gorm"

type Messagem struct {
	ID       uint   `gorm:"primaryKey"`
	Chave    string `gorm:"uniqueIndex"`
	Conteudo string
}

type MessageRepository struct {
	ConDB     *gorm.DB
	Messagens map[string]Messagem
}

func (r *MessageRepository) LoadMessagens() {

	var mensagens []Messagem
	r.ConDB.Find(&mensagens)

	r.Messagens = map[string]Messagem{}
	for _, msg := range mensagens {
		r.Messagens[msg.Chave] = msg
	}

}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{ConDB: db}
}
