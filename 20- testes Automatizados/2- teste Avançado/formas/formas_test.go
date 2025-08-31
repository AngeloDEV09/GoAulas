package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	// TDD - test driven development (Desenvolvimento guiado por testes)
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{Altura: 10, Largura: 15}
		areaEsperada := 150.0
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida é diferente da esperada. Recebida: %f. Esperada: %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida é diferente da esperada. Recebida: %f. Esperada: %f", areaRecebida, areaEsperada)
		}
	})
}
