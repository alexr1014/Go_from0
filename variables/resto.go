package variables

import (
	"fmt"
	"strconv" //paquete dentor de GO para convertir string
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

// mayuscula a la funcion para que sea publica y pueda ser vista fuera del paquete
func RestoVariable() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1600.66
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

// el primer parentesis muestra los valores que requiere/recibe la funcion
// en este caso recibira un numero integer
// el segundo parentesis indica los parametros de salida  si es + de uno deben
//ser separados por una coma

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
