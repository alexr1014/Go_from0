package goroutines

// un goroutine asincrono en GO se usa para ejecutar tareas en paralelo
//para esto cuando se llame la funcion en main se debe anteponer un go
//* (ir a main) y se debe adicionar otra actividad para que las ejecute
import (
	"fmt"
	"strings"
	"time"
)

// recibir el nombre letra por letra
// **en la segunda parte del ejercicio se hablara de channels
func MinombreLento(nombre string, canal1 chan bool) {
	//split toma una cadena y lo separa segun nosotros le ordenemos, en este caso dejar espacio ""
	letras := strings.Split(nombre, "")
	//_ significa que el primer caracter del range no interesa, el segundo es la letra
	for _, letras := range letras {
		//cada segundo se muestra una letra
		//1 seg = 1000 miliseg * time.millisecond
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letras)
	}
	//se le asigna un valor a canal 1 con <-
	canal1 <- true
	//ahora se debe crear el canal dentro de main

}
