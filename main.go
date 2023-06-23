package main

//paquetes fmt permite enviar a la consola texto
//si se importa un solo paquetes usar "" si se importan varios usar ("" "")

//e.g una vez creado el file enteros.go vamos a importar la crpeta variable
//y correr la funcion para calcualr los enteros
import (
	"fmt"
	// nos muestra ambien de ejecucion de donde esta corriendo el sistema
	ejercicio "github.com/alexr1014/Go_from0/ejercicios"
)

// no entiendo porque se requiere declarar las variables estado y texto

func main() {
	/*variables.RestoVariable() - ejercicio no. 1
	estado, texto := variables.ConviertoaTexto(1000) ejercicio no. 2
	fmt.Println(estado)
	fmt.Println(texto)
	condicionales
	if os := runtime.GOOS; os == "linux" || os == "OS X." { //es comparacion operadores &&=and y ||=or
		fmt.Println("Esto no es wimdows, es ", os)
	} else {
		fmt.Println("Esto es wimdows")
	}
	variante condicional switch = multiples valores a evaluar
	switch os := runtime.GOOS; os { // os de esta lina es como la pregunta
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default: // si hay otras opciones se usa defaul (pra el resto)
		fmt.Printf("%s \n", os) //printf se formatera el texto de uan manera
	%s le estoy diciendo que el argumento es un string - \n = salta de linea
	}*/

	// resultado ejericio 01
	num, s := ejercicio.ConvertiraEntero("101")
	fmt.Println(num)
	fmt.Println(s)
}
