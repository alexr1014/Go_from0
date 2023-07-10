package teclado

import (
	"bufio" // lectura por teclado
	"fmt"
	"os" // operative system
	"strconv"
)

var num1 int
var num2 int
var leyenda string
var err error

// la funcion no recibe ni entrega datos, esto se conoce como metodo
func IngresoNumeros() {
	//de donde va leer datos? newscanner el paquete os.stdin (standar input=teclado)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero 1 : ")
	// se ingresa al if si el usuario dio un input con teclado
	if scanner.Scan() {
		//se debe convertir de texto a int (si queremos grabar como variable entero)
		//porque todo l oque entra por bufio entra como texto
		//todo lo que le usuario ingresa entra en la propiedad text
		//como la variable num1 ya fue creada no requiere := los dos puntos
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error()) //si se quiere abortar app panic
		}
	}

	fmt.Println("Ingrese numero 2 : ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error()) //si se quiere abortar app panic
		}
	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Print(leyenda, num1*num2)
}
