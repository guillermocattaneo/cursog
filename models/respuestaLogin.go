package models

/*RespuestaLogin tiene el token que se obtiene en el login*/
type RespuestaLogin struct {
	Token string `json:token, omitempty`
}
