package main

import (
	"fmt"	
	"strconv"
)

func main() {

	//1.Forma tradicional de difinir variables
	var Name string
	var Age int

	Name = "Alice"
	Age = 30

    fmt.Println(Name)
    fmt.Println(Age)

	//2.Definir variable por inferencia de tipo
	Mascota := "Perro"
	Peso := 12.5

	fmt.Println(Mascota)
	fmt.Println(Peso)

	//Conversión de tipos explicita
	numeroCoches := 10
	numeroCarros := float64(numeroCoches)
	numeroVehiculos := int(numeroCarros)
	fmt.Println("numero de vehiculos: ", numeroVehiculos)

	//Convertir string a int
	id := "100"
	i, _ := strconv.Atoi(id)// Ignorar el error con _
	
	fmt.Println(i)
}