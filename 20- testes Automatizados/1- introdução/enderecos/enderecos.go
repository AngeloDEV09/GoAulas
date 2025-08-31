package enderecos

// REGRAS DE NEGÓCIO
import "strings"

//TipoDeEndereco verifica se o endereço começa com um tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmMinusculo := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmMinusculo, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "tipo inválido"
}
