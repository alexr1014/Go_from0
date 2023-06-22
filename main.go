package main

//paquetes fmt permite enviar a la consola texto
//si se importa un solo paquetes usar "" si se importan varios usar ("" "")

//e.g una vez creado el file enteros.go vamos a importar la crpeta variable
//y correr la funcion para calcualr los enteros
import (
	"fmt"

	"github.com/alexr1014/Go_from0/variables"
)

// no entiendo porque se requiere declarar las variables estado y texto

func main() {
	//variables.RestoVariable() - ejercicio no. 1
	estado, texto := variables.ConviertoaTexto(1000)
	fmt.Println(estado)
	fmt.Println(texto)
}
