package models

/*Tweet captura el Bodu del mensaje*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
