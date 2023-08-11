package defer_panic

import (
	"fmt"
	"log"
)

//defer: intruccion que permite configurar cual es la ultima instrucción en ejecutarse cuando salga d la funcion
//uso defer: permite configurar la funcion para cerrar un archivo pase lo que pase cuando la funcion termine

func VemosDefer() {
	fmt.Println("Es es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("este es el segundo mensaje")
}

//panic: aborta programa con un mensje que envia a consola cuando detecta un error o falla

func EjemploPanic() {
	//defer esta acompñado de una funcion anónima
	defer func() {
		reco := recover()
		// if reco= nill significa que hubo un panic
		if reco != nil {
			//log hace lo mismo que un print pero el log antepone la fecha y hora de la ejecución
			//fatalf aborta el programa
			log.Fatalf("ocurrio un error que generó panic \n %v", reco)
		}
		//() al final porqueuna funcion anonima con difer debe terminar asi
	}()
	a := 1
	if a == 1 {
		panic("se encontró el valor 1")
	}
}

//recover: permite recuperarse de un panic
//recover se usa si o si con defer
