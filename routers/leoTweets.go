package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guillermocattaneo/cursog/bd"
)

/*LeoTweets */
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		//http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		http.Error(w, "Debe enviar el parametro ID", http.StatusNotFound) //404 porque no encuentra el tweet
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest) //es un dato requerido
		//http.Error(w, "Debe enviar el parametro pagina", http.StatusNotAcceptable) //O es un NO aceptado?
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if !correcto {
		http.Error(w, "Error al leer los Tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)

}
