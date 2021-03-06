package routers

import (
	"net/http"

	"github.com/guillermocattaneo/cursog/bd"
	"github.com/guillermocattaneo/cursog/models"
)

/*BajaRelacion realiza la baja de la relacion*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		//		http.Error(w, "Ocurrio un error al intentar borrar la relacion. "+err.Error(), http.StatusBadRequest)
		http.Error(w, "Ocurrio un error al dejar de seguir. "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		//http.Error(w, "No se ha logrado borrar la relacion. ", http.StatusBadRequest)
		http.Error(w, "No se ha logrado dejar de seguir. "+err.Error(), http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusOK)
}
