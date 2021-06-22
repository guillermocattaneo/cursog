package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guillermocattaneo/cursog/bd"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina.", http.StatusBadRequest) //otro dato requerido
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a 0.", http.StatusBadRequest) //dato requerido, va con 400
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		//http.Error(w, "Error al leer los tweets.", http.StatusBadRequest)
		http.Error(w, "Error al leer los tweets.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
