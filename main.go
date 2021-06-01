package main

import (
	"log"

	"github.com/guillermocattaneo/cursog/bd"
	"github.com/guillermocattaneo/cursog/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la Base de Datos")
		return
	}
	handlers.Manejadores()
}
