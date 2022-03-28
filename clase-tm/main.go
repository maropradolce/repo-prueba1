package main

import "fmt"

var edad int = 26

func main() {
	//ejercicio1
	nombre := "Mar"
	var direccion = "Buenos Aires"

	fmt.Println(nombre + ", " + direccion)
	fmt.Printf("Nombre: %s, Edad: %d \n", nombre, edad)

	//ejercicio2
	temperatura := 15.7
	humedad := 39.0
	presionAtm := 2.0

	fmt.Printf("La temperatura es %f, la humedad es %f, la presion atmosferica es %f \n ", temperatura, humedad, presionAtm)

	//ejercicio3
	//var 1nombre string
	//var nombre1 string
	//var apellido string
	//var int edad
	//1apellido := 6
	//apellido := 6 // no declarativo el nombre
	//var licencia_de_conducir = true
	//var licenciaDeConducir = true
	//var estatura de la persona int
	//var estaturaDeLaPersona float64
	//cantidadDeHijos := 2

	//ejercicio4
	/*
				var apellido string = "Gomez"
		  		var edad int = "35"
		  		boolean := "false";
		  		var sueldo string = 45857.90
		  		var nombre string = "Julián"
	*/
	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float64 = 45857.90
	var nombre1 string = "Julián"
	fmt.Printf("Mi nombre es %s %s, tengo %d años y mi sueldo es de %f, y soy %t \n", nombre1, apellido, edad, sueldo, boolean)

}
