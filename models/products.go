package models

type Product struct {
	CodProduct  string `json:"product_id" binding:"required"`
	QtdeProduct int64  `json:"product_qtd" binding:"required"`
	VlProduct   int64  `json:"points" binding:"required"`
}

func (p *Product) GetCod() string {
	return p.CodProduct
}

func (p *Product) GetQtde() int64 {
	return p.QtdeProduct
}

func (p *Product) GetValue() int64 {
	return p.VlProduct
}

type ProductPoints interface {
	GetCod() string
	GetQtde() int64
	GetValue() int64
}

type ChatMessage struct {
	Product
}

type Lover struct {
	Product
}

type Presenca struct {
	Product
}

type StreakPresenca struct {
	Product
}

type Churn struct {
	Product
}

type Troca struct {
	Product
}

type JoinFailed struct {
	Product
}

func NewJoinFailed() *JoinFailed {
	return &JoinFailed{
		Product: Product{
			CodProduct:  "41",
			QtdeProduct: 1,
			VlProduct:   100,
		},
	}
}

func NewChatMessage() *ChatMessage {
	return &ChatMessage{
		Product: Product{
			CodProduct:  "5",
			QtdeProduct: 1,
			VlProduct:   1,
		},
	}
}

func NewLover(lover string) *ChatMessage {
	var cod string
	switch lover {
	case "r":
		cod = "13"
	case "airflow":
		cod = "1"
	default:
		cod = ""
	}

	return &ChatMessage{
		Product: Product{
			CodProduct:  cod,
			QtdeProduct: 1,
			VlProduct:   -50,
		},
	}
}

func NewPresent() *Presenca {
	return &Presenca{
		Product: Product{
			CodProduct:  "11",
			QtdeProduct: 1,
			VlProduct:   50,
		},
	}
}

func NewStreakPresent() *StreakPresenca {
	return &StreakPresenca{
		Product: Product{
			CodProduct:  "12",
			QtdeProduct: 1,
			VlProduct:   100,
		},
	}
}

func NewChurn(churn float64) *Churn {
	var cod string
	var points int64
	if churn <= 0.025 {
		cod = "7"
		points = 100
	} else if churn <= 0.05 {
		cod = "8"
		points = 50
	} else if churn <= 0.1 {
		cod = "6"
		points = 10
	} else {
		cod = "38"
		points = 1
	}
	return &Churn{
		Product: Product{
			CodProduct:  cod,
			QtdeProduct: 1,
			VlProduct:   points,
		},
	}
}

func NewTroca() *Troca {
	return &Troca{
		Product: Product{
			CodProduct:  "17",
			QtdeProduct: 1,
			VlProduct:   -1000,
		},
	}
}
