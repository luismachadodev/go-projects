package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a func init")
	n = 10
}

func main() {
	fmt.Println("Func main sendo executada")
	fmt.Println(n)
}
