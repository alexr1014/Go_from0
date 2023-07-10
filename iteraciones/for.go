package iteraciones

import (
	"fmt"
)

// for sintaxis
func Iterar() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	/* otra manera de escribir el codigo
	for i := 0; i < 10; i++ {
		si quiero que i incremente en cierto numero -> i+=5 o i-=5 para disminuir
		fmt.Println(i)
	}*/

}
