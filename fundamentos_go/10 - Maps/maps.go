package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Luis",
		"sobrenome": "Henrique",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Dora",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Direito",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Escorpiao",
	}

	fmt.Println(usuario2)
}
