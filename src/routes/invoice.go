package routes

import (
	"encoding/json"
	"fmt"

	"example.com/taller/src/models"
	"example.com/taller/src/structures"
	"github.com/savsgio/atreugo/v11"
)

func GetInvoice(ctx *atreugo.RequestCtx) error {
	response, statusCode := models.GetMedicines()
	return ctx.JSONResponse(response, statusCode)
}

func CreateInvoice(ctx *atreugo.RequestCtx) error {
	var newInvoice structures.Invoice
	err := json.Unmarshal(ctx.PostBody(), &newInvoice)
	if err != nil {
		fmt.Println(err)
	}
	message, statusCode := models.CreateInvoice(newInvoice.FechaCreacion, newInvoice.PagoTotal, newInvoice.IdPromocion, newInvoice.IdMedicamentos)
	return ctx.TextResponse(message, statusCode)
}
