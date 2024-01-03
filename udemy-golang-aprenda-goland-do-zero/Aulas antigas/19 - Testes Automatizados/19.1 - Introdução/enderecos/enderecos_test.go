package enderecos

// TESTE DE UNIDADE
// É um teste que testa a menor unidade do meu código
// No caso, será uma função

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// Sempre começa com a palavra Test
	// Seguido do nome da função

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"RUA ABC", "Tipo inválido"},
		{"Pç da Bandeira", "Tipo inválido"},
		{"Pça da Imprensa", "Tipo inválido"},
		{"Av. Suburbana", "Tipo inválido"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
