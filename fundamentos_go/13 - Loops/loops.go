package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Luis", "Joao", "Davi"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Luis",
		"sobrenome": "Henrique",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
