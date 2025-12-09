package errors

import "errors"

var (
	ErrPresencaAssinadaAnterior = errors.New("lista já foi assinada hoje")
	ErrUsuarioNaoCriado         = errors.New("usuario ainda não foi criado")
	ErrUsuarioExistente         = errors.New("usuario já existente")
	ErrRequisicaoNaoOK          = errors.New("erro na requisicao HTTP")
	ErrPontosInsuficientes      = errors.New("pontos insuficientes para troca")
	ErrClasseNaoEncontrada      = errors.New("classe não encontrada")
	ErrRacaNaoEncontrada        = errors.New("raça não encontrada")
	ErrLootJaResgatado          = errors.New("loot ja foi resgatado hoje")
	ErrItemNaoEncontrado        = errors.New("item nao encontrado")
)
