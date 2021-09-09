package main

import (
	"encoding/json"
	"fmt"

	"example.com/taller/src/database"
	"example.com/taller/src/models"
	_ "github.com/lib/pq"
	"github.com/savsgio/atreugo"
)

func main() {
	config := &atreugo.Config{
		Host: "localhost",
		Port: 8000,
	}
	server := atreugo.New(config)

	database.GetConnection()

	server.Path("POST", "/medicamento", func(ctx *atreugo.RequestCtx) error {
		var newMedicine models.Medicine
		err := json.Unmarshal(ctx.PostBody(), &newMedicine)
		if err != nil {
			fmt.Println(err)
		}
		return models.InsertMedicine(newMedicine.Nombre, newMedicine.Precio, newMedicine.Ubicacion, ctx)
	})

	server.Path("GET", "/medicamento", func(ctx *atreugo.RequestCtx) error {
		return models.GetMedicines(ctx)
	})

	// promociones
	server.Path("POST", "/promocion", func(ctx *atreugo.RequestCtx) error {
		var newPromotion models.Promotion
		err := json.Unmarshal(ctx.PostBody(), &newPromotion)
		if err != nil {
			fmt.Println(err)
		}
		return models.InsertPromotion(newPromotion.Descripcion, newPromotion.Porcentaje, newPromotion.FechaInicio, newPromotion.FechaFin, ctx)
	})

	server.Path("GET", "/promocion", func(ctx *atreugo.RequestCtx) error {
		return models.GetPromotions(ctx)
	})

	// Facturas
	server.Path("POST", "/factura", func(ctx *atreugo.RequestCtx) error {
		var invoice models.Invoice
		err := json.Unmarshal(ctx.PostBody(), &invoice)
		if err != nil {
			fmt.Println(err)
		}
		return models.InsertInvoice(invoice.FechaCreacion, invoice.PagoTotal, invoice.IdPromocion, invoice.IdMedicamentos, ctx)
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
