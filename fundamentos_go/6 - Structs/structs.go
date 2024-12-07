package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Luis"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Tal", 200}

	usuario2 := usuario{"Luis Henrique", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Giovana"}
	fmt.Println(usuario3)
}
