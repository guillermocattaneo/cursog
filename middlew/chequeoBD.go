package middlew

import (
	"net/http"

	"github.com/guillermocattaneo/cursog/bd"
)

/*ChequeoBD es el middleWare para conocer el estado de la BdD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
