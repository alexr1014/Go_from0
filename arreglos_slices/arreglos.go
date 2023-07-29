package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{45, 8, 7} //genera un arreglo llamado tabla de 10 elementos de tipo int
//luego del = es la menra de asignar valores al arreglo
var matriz [20][30]int // la matriz tiene 20filas por 30 columnas (o 30 elementos)
func MuestroArreglos() {
	tabla[2] = 34
	tabla[7] = 87

	//variacion del arreglo
	tabla2 := [10]string{"Juan", "Pepe", "Alex"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[1][2] = 15 // le decimos que asigne el valor 15 en el vector 1 y la posicion 2 del vector 1
	fmt.Println(matriz)
}

// matriz es un vector de vectores
