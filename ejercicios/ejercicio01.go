package ejercicio

import (
	"strconv"
)

func ConvertiraEntero(s string) (int, string) {
	// Atoi tiene 2 salidas (int,error) e.g se usa "num, _" (_ si solo interesa el numero se usa guion bajo
	// significa que el argumento de resultado no se asignara a ninguna variable)
	// pero si se usa num, err el error sera asignado a la variable "err"
	// nil = null en GO
	// err,Error devuelve el error
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
