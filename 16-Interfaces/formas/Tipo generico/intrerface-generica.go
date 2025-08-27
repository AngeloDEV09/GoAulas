package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Interfaces como tipo Genérico")

	generica("String")
	generica(1)
	generica(true)

	fmt.Println(1, 2, 3, "string", false, true, float64(12345))

	// NÃO FAZER GAMBIARRA DESSE TIPO.
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)

}
