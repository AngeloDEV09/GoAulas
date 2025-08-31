// TESTE DE UNIDADE
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// TestTipoDeEndereco verifica se a função TipoDeEndereco está funcionando corretamente
func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{enderecoInserido: "Rua ABC", retornoEsperado: "rua"},
		{enderecoInserido: "Avenida Paulista", retornoEsperado: "avenida"},
		{enderecoInserido: "Estrada do Sol", retornoEsperado: "estrada"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço recebido é diferente do esperado. Recebido: %s. Esperado: %s.",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
