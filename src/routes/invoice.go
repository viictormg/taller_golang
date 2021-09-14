package routes

import (
	"encoding/json"
	"fmt"
	"strings"

	"example.com/taller/src/models"
	"example.com/taller/src/structures"
	"github.com/savsgio/atreugo/v11"
)

func GetInvoice(ctx *atreugo.RequestCtx) error {
	fechaInicial := string(ctx.QueryArgs().Peek("fecha_inicio"))
	fechaFinal := string(ctx.QueryArgs().Peek("fecha_fin"))
	response, statusCode := models.GetInvoice(fechaInicial, fechaFinal)

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

func SimulateInvoice(ctx *atreugo.RequestCtx) error {
	fecha := string(ctx.QueryArgs().Peek("fecha"))
	IdMedicamentosString := string(ctx.QueryArgs().Peek("id_medicamentos"))
	IdMedicamentos := strings.Split(IdMedicamentosString, ",")

	descuento, statusCode := models.SimulateInvoice(fecha, IdMedicamentos)
	return ctx.TextResponse(fmt.Sprintf("%v", descuento), statusCode)
}
