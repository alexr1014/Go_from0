package ejer_interfaces

//paso 3
// crear nueva funcion
import (
	"fmt"

	"github.com/alexr1014/Go_from0/interfaces"
)

// en el paso 4 se creo el archivo mujer
// cuando una estructura implementa las mismas funciones y metodos d euna interfaz,
// automaticamente se convierte en ese objeto. En este caso tanto Hombre como Mujer por implementar
// lo mismo que la interfaz Humano, van hacer de tipo humano
func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
