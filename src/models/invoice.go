package models

import (
	"fmt"
	"net/http"
	"strconv"

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
	var IdMedicamentosString string
	IdMedicamentosString = "array["

	for i := 0; i < len(IdMedicamentos); i++ {
		coma := ","
		if i == len(IdMedicamentos)-1 {
			coma = ""
		}
		IdMedicamentosString += strconv.Itoa(IdMedicamentos[i]) + coma
	}
	IdMedicamentosString += "]"

	query := fmt.Sprintf("insert into factura (fecha_creacion, pago_total, id_promocion, id_medicamentos) values ('%s', '%f', '%v', %s) returning id", FechaCreacion, PagoTotal, IdPromocion, IdMedicamentosString)

	fmt.Println(query)
	db.QueryRow(query).Scan(&invoiceId)

	if invoiceId == 0 {
		return ctx.TextResponse("medicamento no creado", http.StatusBadRequest)

	}
	return ctx.TextResponse("medicamento creado", http.StatusOK)
}
