package interfaces

//Paso 5. se crea la rama Servivo, la cual es el parent object de las interfaces humano, animaly vegetal
//todas las interfaces estas amarradas a Servivio cuando añadimos el metodo Estavivo() en la estructura de cada interface

type Servivo interface {
	Estavivo() bool
}
