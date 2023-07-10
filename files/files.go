package files

//se crea en conjunto con el ejercicio dos, despues de tener el string concatenado
//la idea es guardar el ejercicio en este archivo

import (
	"bufio" //capturar datos
	"fmt"
	"os"

	ejercicio "github.com/alexr1014/Go_from0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

// funcion GrabaTabla: toma tabla de multiplicar y genera archivo .txt
func GrabaTabla() {
	var err error
	var texto string = ejercicio.TablaMultiplicar()
	archivo, err := os.Create(filename) // .Create crea archivo en el path folder txt construido (var filenames tiene asignafo el path)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return //si existe error el return terminar el metodo o funcion
	} else {
		//se guardara la informacion (texto) en la variable archivo
		fmt.Fprintln(archivo, texto)
	}
	archivo.Close() //todo archivo que se abre debe cerrarse
}

// funcion SumaTabla: va a tomar el archivo y agrupar diferentes tablas de multiplicar
func SumaTabla() {
	var texto string = ejercicio.TablaMultiplicar()
	if Append(filename, texto) == false {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}
	//no es necesario el else porque se tiene return
	//la siguiente parte es si append se jecuta que pasa
	//writeString se usa para grabar el texto
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true
}

// se comienza con escritura y luego lectura
// func LeoArchivo muestra el resultado en consola
func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	// esta es la parte verdadera (true)
	// bufio toma los datos de archivo
	scanner := bufio.NewScanner(archivo)
	//me devuelve cada una de las lienas del archivo
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
}
