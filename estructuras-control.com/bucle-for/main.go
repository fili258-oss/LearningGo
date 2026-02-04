package main

import(
	"fmt"
)

func main() {
	//1.forma clasica con un incrementador
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//2.forma de while
	i:= 0
	for i < 5 {
		i++
		if i < 3 {
			continue//saltarse una interacción
		}
		break// salir de bucle
	}
}

