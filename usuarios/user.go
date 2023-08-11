package usuarios

import (
	"fmt"
	"time"

	"github.com/alexr1014/Go_from0/modelos"
)

func AltaUsuario() {
	//el objeto se crea con el comando new
	u := new(modelos.User) // u es un objeto creado para llamar la estructura en el package modelos
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
