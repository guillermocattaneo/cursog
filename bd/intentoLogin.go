package bd

import (
	"github.com/guillermocattaneo/cursog/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin compara los passwords */
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passwordByte := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordByte)
	if err != nil {
		return usu, false

	}
	return usu, true
}
