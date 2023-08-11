package modelos

//en esta parte se crean en conjunto los package modelos y usuarios que seran usados conjuntamente
import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	status    bool
}

// entre func y el nombre de la funcion poner la estructura que se quiere trabajar
// entre los elementos de la funcion aparecen los elementos de la estructura
// * es un puntero de memoria que indica donde GO tiene creada esa estructura
// *User significa que la funcion AddUser se relaciona con la estructura User
func (this *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.status = status

}
