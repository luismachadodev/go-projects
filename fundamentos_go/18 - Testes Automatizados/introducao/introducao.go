package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia Tal")
	fmt.Println(tipoEndereco)
}
