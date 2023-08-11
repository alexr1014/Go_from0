package main

//paso 5. llamar el package ejer_interfaces (la e al inicio )
//se puede agregar una e al comienzo de import, y esta e reemplaza el nombre
//"ejer_interfaces para evitar su longitud al llamar llamar la funcion"
import (
	"github.com/alexr1014/Go_from0/middlewares"
)

//paquetes fmt permite enviar a la consola texto
//si se importa un solo paquetes usar "" si se importan varios usar ("" "")

//e.g una vez creado el file enteros.go vamos a importar la crpeta variable
//y correr la funcion para calcualr los enteros

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
	default: // si hay otras opciones se usa defaul (para el resto)
		fmt.Printf("%s \n", os) //printf se formatera el texto de uan manera
	%s le estoy diciendo que el argumento es un string - \n = salta de linea
	}

	// resultado ejericio 01
	num, s := ejercicio.ConvertiraEntero("101")
	fmt.Println(num)
	fmt.Println(s)

	teclado.IngresoNumeros()

	iteraciones.Iterar()*/

	//fmt.Println(ejercicio.TablaMultiplicar())

	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.IncrementoTemp()

	//e reemplaza el nomrbe largo de ejer_interfaces
	//Aqui aplica el concepto de propiedades iguales y asociados a la interface humano
	/*Pedro := new(modelos.Hombre)
	e.HumanoRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanoRespirando(Maria)*/

	//canal1 := make(chan bool)
	//*goroutine
	//go goroutines.MinombreLento("Iniesta", canal1)
	//**como se le dice detengase hasta que no se termine la groutine

	//defer func() {
	//	<-canal1 //espere a que canal 1 termine la ejecucion
	//}()
	//fmt.Println("Estoy aqui")
	/* con el chanel esta parte ya no es neesario
	var x string
	fmt.Scanln(&x)*/

	middlewares.MiMiddleware()
}
