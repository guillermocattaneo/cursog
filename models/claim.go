package models

import (
	"github.com/dgrijalva/jwt-go"
	//jwt "github.com/guillermocattaneo/cursog/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim es la estructura para procesar el JWT*/
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
