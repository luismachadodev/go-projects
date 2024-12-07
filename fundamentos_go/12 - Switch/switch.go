package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Numero Invalido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Domingooo"
	case numero == 3:
		diaDaSemana = "Domingo"
	case numero == 4:
		diaDaSemana = "Domingo"
	case numero == 5:
		diaDaSemana = "Domingo"
	case numero == 6:
		diaDaSemana = "Domingo"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Numero Invalido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	fmt.Println("---------")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
