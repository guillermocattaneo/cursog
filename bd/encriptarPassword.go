package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword hace pizzas*/
func EncriptarPassword(pass string) (string, error) {
	//cantidad de veces que pasara sobre el texto para encriptarlo, se aconseja de 6 a 8
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err

}
