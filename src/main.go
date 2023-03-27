package main

import "fmt"

func main() {

	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaracion de variables enteras
	base := 12          //esos : indica que esta variable no ha sido declarada anterioremnte y una vez que := va a crear esa varable Y le asigna ese valor parseando el tipo de dato
	var altura int = 14 //indico que es una variable, la nombre, le indico el tipo de dato y le asigno el valor. Es explicito
	var area int        // todo lo anterior co la diferencia de que no asigo el valor

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

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x // = porque la variable ya la tenemos definida. Pero que pasa si el tipo de variable se modifica con la nueva cuenta?
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Decremental", x)

	//Retos
	// Area del rectangulo, del trapecio y de un circulo
	const baseRectangulo = 5
	const alturaRectangulo = 3.5
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area del rectangulo", areaRectangulo)

	const minorbaseTrapecio = 5
	const mayorBaseTrapecio = 3.25
	const alturaTrapecio = 2.1
	areaTrapecio := (minorbaseTrapecio + mayorBaseTrapecio) / 2 * alturaTrapecio
	fmt.Println("Area del trapecio", areaTrapecio)

	const PI float64 = 3.141516
	const radioCirculo = 2.6
	areaCirculo := PI * radioCirculo * radioCirculo
	fmt.Println("Area del Circulo", areaCirculo)
}
