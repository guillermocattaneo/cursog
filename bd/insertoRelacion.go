package bd

import (
	"context"
	"time"

	"github.com/guillermocattaneo/cursog/models"
)

/*InsertoRelacion de usuarios en la BdD*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("curso")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
