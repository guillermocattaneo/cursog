package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/guillermocattaneo/cursog/bd"
	"github.com/guillermocattaneo/cursog/models"
)

/*SubirAvatar*/
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		//http.Error(w, "Error al grabar el avatar en la BdD! "+err.Error(), http.StatusBadRequest)
		http.Error(w, "Error al grabar el avatar en la BdD! "+err.Error(), http.StatusGatewayTimeout)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
