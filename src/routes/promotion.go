package routes

import (
	"encoding/json"
	"fmt"

	"example.com/taller/src/models"
	"example.com/taller/src/structures"
	"github.com/savsgio/atreugo/v11"
)

func CreatePromotion(ctx *atreugo.RequestCtx) error {
	var newPromotion structures.Promotion
	err := json.Unmarshal(ctx.PostBody(), &newPromotion)
	if err != nil {
		fmt.Println(err)
	}
	message, statusCode := models.CreatePromotion(newPromotion.Descripcion, newPromotion.Porcentaje, newPromotion.FechaInicio, newPromotion.FechaFin)

	return ctx.TextResponse(message, statusCode)
}

func GetPromotions(ctx *atreugo.RequestCtx) error {
	response, statusCode := models.GetPromotions()
	return ctx.JSONResponse(response, statusCode)
}
