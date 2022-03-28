/*
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.
Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/

package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	//A
	fmt.Println(employees["Benjamin"])
	//B
	for key, employeeAge := range employees {
		if employeeAge > 21 {
			fmt.Printf("%s es mayor de 21 años \n", key)
		}
	}
	//C
	employees["Federico"] = 25
	//D
	delete(employees, "Pedro")
	fmt.Println(employees)
}
