package main

import "fmt"

func main() {

	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaracion de variables enteras
	base := 12 //esos : indica que esta variable no ha sido declarada anterioremnte y una vez que := va a crear esa varable Y le asigna ese valor parseando el tipo de dato
	var altura int = 14 //indico que es una variable, la nombre, le indico el tipo de dato y le asigno el valor. Es explicito
	var area int // todo lo anterior co la diferencia de que no asigo el valor

	fmt.Println(base, altura, area)

	// Zero values: valor que va a tener por defecto al momento de 
	// crear una variable y no le asignamos un valor 
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)
}
