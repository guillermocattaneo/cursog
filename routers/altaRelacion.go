package routers

import (
	"net/http"

	"github.com/guillermocattaneo/cursog/bd"
	"github.com/guillermocattaneo/cursog/models"
)

/*AltaRelacion registra la relacion entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		//http.Error(w, "Ocurrio un error al intentar insertar la relacion. "+err.Error(), http.StatusBadRequest)
		http.Error(w, "Ocurrio un error al intentar seguir. "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		//http.Error(w, "No se ha logrado insertar la relacion. ", http.StatusBadRequest)
		http.Error(w, "No se ha logrado seguir. ", http.StatusInternalServerError) //500
		return
	}

	//Como lo tomo como una creacion de relacion entre usuarios, lo "creamos" y no mandamos un simple 200
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusCreated)
}
