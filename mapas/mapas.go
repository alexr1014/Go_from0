package mapas

import "fmt"

//los mapas son conexiones de elementos
//mapa es un coleccion de clave-valor, tiene una clave y la clave esta asociada un valor

func MostrarMapas() {
	//hay dos manera de generar estructuras de mapas
	//1. por asignacion usando los elementos make y map
	paises := make(map[string]string) //la clave y el valor son de tipo string

	paises["Mexico"] = "D.F"
	paises["Colombia"] = "Bogotá"
	paises["Brasil"] = "Brasilia"
	paises["Portugal"] = "Lisboa"
	fmt.Println(paises)
	fmt.Println(paises["Colombia"])
	//2. Por asignacion directa usanso el elemento map

	campeonato := map[string]int{
		"Millonarios": 25,
		"Nacional":    30,
		"Deportivo":   45,
		"Neiva":       21,
	}

	fmt.Println(campeonato)

	//a diferencia de vectores y slices, cuando se consulta un mapa con "range" devuelve clave y el valor
	//al for se asingan las dos variables directamente

	for equipo, puntaje := range campeonato {

		fmt.Printf("\nEl equipo %s alcanzó un puntaje de %d", equipo, puntaje)
	}
	//como eliminar un elemento de mapas
	delete(campeonato, "Deportivo")
	fmt.Println(campeonato)

	//como consultar el valor de un equipo

	puntaje, existe := campeonato["Millonarios"]

	fmt.Printf("El puntaje del equipo es %d y el equipo existe? %t", puntaje, existe)
}
