package models

import (
	"fmt"
	"net/http"

	"example.com/taller/src/database"
	"github.com/savsgio/atreugo"
)

type Invoice struct {
	Id             int     `json:"id"`
	FechaCreacion  string  `json:"fecha_creacion"`
	PagoTotal      float32 `json:"pago_total"`
	IdPromocion    int     `json:"id_promocion"`
	IdMedicamentos []int   `json:"id_medicamentos"`
}

func InsertInvoice(FechaCreacion string, PagoTotal float32, IdPromocion int, IdMedicamentos []int, ctx *atreugo.RequestCtx) error {
	db := database.GetConnection()
	var invoiceId int
	query := fmt.Sprintf("insert into factura (fecha_creacion, pago_total, id_promocion) values ('%s', '%f', '%v')", FechaCreacion, PagoTotal, IdPromocion)

	// query := fmt.Sprintf("insert into medicamento(nombre, precio, ubicacion) values ('%s', %f, '%s') returning id", nombre, precio, ubicacion)
	fmt.Println(query)
	db.QueryRow(query).Scan(&invoiceId)

	if invoiceId == 0 {
		return ctx.TextResponse("medicamento no creado", http.StatusBadRequest)

	}
	return ctx.TextResponse("medicamento creado", http.StatusOK)
}
