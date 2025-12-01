package repositories

import "math/rand/v2"

type MensagemAleatoria struct {
	Name    string
	Options []string
}

func (m *MensagemAleatoria) GetMensagem() string {
	n := len(m.Options)
	index := rand.IntN(n)
	return m.Options[index]
}

var mensagemCoach = MensagemAleatoria{
	Name: "coach",
	Options: []string{
		"%s, no tempo certo, tudo dará errado.",
		"%s, acorda, o fracasso te espera.",
		"%s, nunca subestime sua incapacidade.",
		"%s, seja forte, desista!",
		"%s, fracasse enquanto eles descansam.",
		"%s, a vida te derruba hj, preparando para a queda de amanhã.",
		"%s, nunca é tarde para desistir.",
		"%s, você é mais fraco do que pensa.",
		"%s, nada vem no tempo certo.",
		"%s, todo fracasso começa com a decisão de tentar.",
		"%s, você só perderá amanhã se não desistir hoje.",
		"%s, desistir é sempre a melhor opção.",
		"%s, descubra novas formas de fracassar.",
		"%s, a expectativa é mãe da merda!",
		"%s, nunca aceite críticas construitivas de quem nunca construiu nada.",
		"%s, continue sonhando, você acabará se atrasando",
	},
}

type MensagensAleatoriasRepository struct {
	Mensagens map[string]MensagemAleatoria
}

func (r *MensagensAleatoriasRepository) LoadMensagensAleatorias() {

	r.Mensagens = map[string]MensagemAleatoria{
		mensagemCoach.Name: mensagemCoach,
	}

}

func NewMensagensAleatoriasRepository() *MensagensAleatoriasRepository {
	repo := &MensagensAleatoriasRepository{}
	repo.LoadMensagensAleatorias()
	return repo
}
