package modelos

import "fmt"

type Joven struct {
	salario int
}

func (j *Joven) Respirar()    { fmt.Printf("Hola") }
func (j *Joven) Pensar()      { fmt.Printf("Soy un joven y estoy pensando") }
func (j *Joven) Comer()       { fmt.Printf("Soy un joven y estoy comiendo") }
func (j *Joven) Sexo() string { return "Joven" }
