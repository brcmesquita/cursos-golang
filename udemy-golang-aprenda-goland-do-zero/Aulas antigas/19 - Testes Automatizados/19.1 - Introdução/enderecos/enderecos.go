package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia", "beco", "travessa", "praça"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	// Esse pacote strings.Split, irá dividir a string a cada espaço em branco
	// Exemplo:
	// O usuário adiciona "Rua dos Bobos"
	// O Split divide (faz o split) e iprime "Rua", "dos", "Bobos"
	// Como definimos a posição que queremos no final do "primeiraPalavraDoEndereco"
	// ele irá retornar a primeira posição do Slice -> [0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
