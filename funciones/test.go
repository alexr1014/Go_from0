package funciones

import "fmt"

func HaceCalor(temp float64) string {
	if temp < 20 {
		fmt.Println("la temperatura exterior es de", temp, "°C")
		return ("Hace frio")

	} else {
		fmt.Println("la temperatura exterior es de", temp, "°C")
		return ("Hace calor")
	}

}

func IncrementoTemp() {

	for temp := 0.0; temp <= 30; temp += 5 {
		x := HaceCalor(temp)
		fmt.Println(x)
	}
}
