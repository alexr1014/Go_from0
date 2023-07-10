package ejercicio

/* despues de terminar el ejercicio 2 se genera una variacion y es guaradar
todo el ejercicio en una variable y la devuelva por parametrod e respuesta de
funcion. Para ue una vez obtenida ese string se pueda grabar como un archivo*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {

	var err error
	var numero int
	var texto string //se crea la variable texto, ahi vamos a guardar el ejercicio

	ingreso := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero: ")
		if ingreso.Scan() {
			numero, err = strconv.Atoi(ingreso.Text())
			if err != nil {
				continue
			} else {
				break
			}

		}
	}

	i := 1
	fmt.Println("la tabla del numero", numero)
	for i <= 10 {
		//guardar esta seccion como string, se hace con Sprintf (devuelve string)
		// y concatenar con una funcion texto
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
		i += 1

	}
	return texto
}
