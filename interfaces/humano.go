package interfaces

//paso 1
// una interfaz es modelos de comportamiento que se pueden asociar a las estructuras
// aqui no se colocan variables ni tipos de datos
type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
	Estavivo() bool
}
