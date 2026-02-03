package main

import (
	"fmt"

)

func main(){
	fmt.Println("Operadores en Go")

	//Operadores Aritmeticos

	a := 10
	b := 3

	suma := a + b
	resta := a - b
	multiplicacion := a * b
	division := a / b
	residuo := a % b

	fmt.Println("Suma: ", suma)
	fmt.Println("Resta: ", resta)
	fmt.Println("Multiplicacion: ", multiplicacion)
	fmt.Println("Division: ", division)
	fmt.Println("Modulo: ", residuo)

	//Operadores Lógicos
	c := true
	d := false

	fmt.Println(c && d)
	fmt.Println(c || d)
	fmt.Println(!c)

	//Valores condicionales
	e := 5
	f := 10
	fmt.Println("---Operadores condicionales---")
	fmt.Println(e == f) //output: false
	fmt.Println(e != f) //output: true
	fmt.Println(e < f) //output: true
	fmt.Println(e > f) //output: false
	fmt.Println(e >= f) //output: false
	fmt.Println(e <= f) //output: true
}