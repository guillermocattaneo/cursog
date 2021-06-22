package middlew

import (
	"net/http"

	"github.com/guillermocattaneo/cursog/routers"
)

/*ValidoJWT valida el JWT que viene de la peticion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			//http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
