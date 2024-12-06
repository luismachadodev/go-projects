package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 1000000
	fmt.Println(numero)

	var numero2 uint32 = 1000 // unsigned int
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 12356
	fmt.Println(numero3)

	// int8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	var str string = "Qualquer texto"
	fmt.Println(str)

	str2 := "Outro texto"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
