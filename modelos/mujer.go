package modelos

//paso 4. hacer l omismo para la mujer
//para copiar todo "herencia de propiedades del hombre"

type Mujer struct {
	Hombre
}

func (h *Mujer) Respirar()      { h.respirando = true }
func (h *Mujer) Pensar()        { h.pensando = true }
func (h *Mujer) Comer()         { h.comiendo = true }
func (h *Mujer) Sexo() string   { return "Mujer" } //sexo es un string
func (h *Mujer) Estavivo() bool { return h.estavivo == true }
