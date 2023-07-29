package arreglos_slices

// un slice es un arreglo dinamico
import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{34, 78, 9, 56, 34, 78, 12, 8, 9}

func MuestroSlice() {

	fmt.Println(tablaS)

	arreglo1 := arreglo[3:]
	arreglo2 := arreglo[:7]
	arreglo3 := arreglo[5:9]

	fmt.Println(arreglo1)
	fmt.Println(arreglo2)
	fmt.Println(arreglo3)

}

//slice maneja dos valores: dimension(cantidad de elemtnos) y capacidad (espacio en memoria que go
//reserva para crear un slice que no es la cantidad de elementos)
//GO reserva mas memoria para mejorar performance
//slice son dinamicos. Esto significa que se puede adicionar capacidad dependiendo de los datos que lleguen

func Capacidad() {
	//la instruccion make permite crear un slice vacio y darle una capcidad
	elementos := make([]int, 5, 20) //se crea un slice vacio [] con 5 elementos iniciales y capacidad para 20
	fmt.Printf("largo %d, capacidad %d", len(elementos), cap(elementos))
	//cap es una funcion para obtener la capcidad de un slice
	//De esta manera GO va tarde muy poco tiempo los elementos porque ya resefvo la capacidad

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)

	}
	fmt.Println(nums)
	fmt.Printf("\nlargo %d, capacidad %d", len(nums), cap(nums))
	//el resultado arroja un largo de 100 y capacidad de 128, este ultimo porque
	//a medida que crece el slice GO reserva mas espacio
}
