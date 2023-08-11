package modelos

//paso 2
// este archivo esta relacionado con el archivo humano.go de interfaces

type Hombre struct {
	edad       int
	altura     float64
	peso       float64
	respirando bool
	pensando   bool
	comiendo   bool
	estavivo   bool
}

// para que la struct Hombre implemente la interfaces Humano, Hombre debe
// tener las mismas funbciones que humano

// * significa que la funcion pertenece a la estructura Hombre
//Respirar en mayuscula porque asi esta en la interface

func (h *Hombre) Respirar()      { h.respirando = true }
func (h *Hombre) Pensar()        { h.pensando = true }
func (h *Hombre) Comer()         { h.comiendo = true }
func (h *Hombre) Sexo() string   { return "Hombre" } //sexo es un string
func (h *Hombre) Estavivo() bool { return h.estavivo == true }
