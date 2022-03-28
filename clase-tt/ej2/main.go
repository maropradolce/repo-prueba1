/*
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

import "fmt"

func main() {
	var (
		age               = 45
		isEmployed        = true
		seniorityInMonths = 15
		salary            = 100100.0
	)
	if age > 22 && isEmployed == true && seniorityInMonths > 12 {
		if salary > 100000.0 {
			fmt.Println("Puede recibir préstamos y NO recibirá intereses")
		} else {
			fmt.Println("Puede recibir préstamos y SI recibirá intereses")
		}
	} else {
		fmt.Println("No puede recibir préstamos")
	}

}
