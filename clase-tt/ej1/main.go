/*
Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras.
*/

package main

import "fmt"

func main() {
	palabra := "gato"

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c \n", palabra[i])
	}

	fmt.Println(len(palabra))
}
