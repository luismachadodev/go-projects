package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Luis", "Machado", 20, 166}
	fmt.Println(p1)

	e1 := estudante{p1, "Sistema de Informações", "UFN"}
	fmt.Println(e1)
}
