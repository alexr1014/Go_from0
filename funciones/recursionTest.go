package funciones

import "fmt"

func RecInversa(n float64) {

	if n <= 10.0 {
		return
	}
	fmt.Println(n)
	RecInversa(n / 100)
}
