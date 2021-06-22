package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweetsSeguidores es la estructura para devolver los tweets*/
type DevuelvoTweetsSeguidores struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID string             `bson:"usuarioid" json:"userid,omitempty"`
	//UsuarioID         string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string `bson:"usuarioRelacionid" json:"userRelationId,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
