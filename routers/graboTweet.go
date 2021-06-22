package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/guillermocattaneo/cursog/bd"
	"github.com/guillermocattaneo/cursog/models"
)

/*GraboTweet graba el tweet en la BdD */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		//http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente "+err.Error(), 400)
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		//http.Error(w, "No se ha logrado insertar el Tweet ", 400)
		http.Error(w, "No se ha logrado insertar el Tweet ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) //Se deja el Creado porque se CREA el mensaje como registro.

}
