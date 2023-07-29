package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor     // 2
	secuencia := 0      // 0, 1, 2
	return func() int { //la funcion retorno comprende las filas 8 a 12
		secuencia++ // 1, 2, 3
		fmt.Println("esta es la secuencia", secuencia)
		return numero * secuencia // 2, 4, 6
	}
}

func LlamarClosure() {
	tabladel := 2
	Mitabla := tabla(tabladel) //en este punto aun no se ejecuta la funcion tabla, la variable Mitabla se convierte en funcion al asignarle una funcion
	for i := 1; i <= 10; i++ {
		fmt.Println("este es i", i)
		fmt.Println(Mitabla()) //aqui se ejecuta la funcion Mitabla(), iniciamos con el 2
	}
}
