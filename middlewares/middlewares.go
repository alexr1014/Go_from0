package middlewares

//funcion previa a funciones para que hagan cierta tarea
//e.g puede ejecutar varias funciones a partir de una funcion "madre"
//un middleware debe tener el mismo tipo de datos siempre

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")
	//a la nueva funcion "operacionesMidd" se le asigna cada una de las funciones creadas
	//todas las funciones se llaman desd euna misma funcion
	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(10, 6)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 4)
	fmt.Println(result)
}

//crear ahora la funcion middleware -> operacionesMidd
//el middleware debe recibir los mismo parametros de las funcioens suma, resta, etc
//y va a devolver lo mismo, entonces repite los datos
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		//este retorno hace referencia a la funcion sumar, restar...
		return f(a, b)
	}
}
