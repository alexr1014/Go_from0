package variables

//crear funcionq ue muestre por consola tres variables enteras

import "fmt"

//para que una variable o funcion sea llamada desde fuera del paquetes tiene que iniciar en mayuscula.
//en Go las variables se pueden declarara de dos maneras: declarativas o por asignacion
func MuestroEnteros() {
	var intcomun int      // declarativa
	intde32 := int32(10)  // asignacion
	intde64 := int64(100) //asignacion

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
