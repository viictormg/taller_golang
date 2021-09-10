package routes

import (
	"encoding/json"
	"fmt"

	"example.com/taller/src/models"
	"example.com/taller/src/structures"
	"github.com/savsgio/atreugo/v11"
)

func GetMedicine(ctx *atreugo.RequestCtx) error {
	response, statusCode := models.GetMedicines()
	return ctx.JSONResponse(response, statusCode)
}

func CreateMedicine(ctx *atreugo.RequestCtx) error {
	var newMedicine structures.Medicine
	err := json.Unmarshal(ctx.PostBody(), &newMedicine)
	if err != nil {
		fmt.Println(err)
	}
	message, statusCode := models.CreatetMedicine(newMedicine.Nombre, newMedicine.Precio, newMedicine.Ubicacion)
	return ctx.TextResponse(message, statusCode)
}
