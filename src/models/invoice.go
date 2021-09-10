package models

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/taller/src/database"
)

type Invoice struct {
	Id             int     `json:"id"`
	FechaCreacion  string  `json:"fecha_creacion"`
	PagoTotal      float32 `json:"pago_total"`
	IdPromocion    int     `json:"id_promocion"`
	IdMedicamentos []int   `json:"id_medicamentos"`
}

func CreateInvoice(FechaCreacion string, PagoTotal float32, IdPromocion int, IdMedicamentos []int) (string, int) {
	db := database.GetConnection()
	var invoiceId int
	query := fmt.Sprintf("insert into factura (fecha_creacion, pago_total, id_promocion) values ('%s', '%f', '%v') returning id", FechaCreacion, PagoTotal, IdPromocion)

	db.QueryRow(query).Scan(&invoiceId)

	if invoiceId == 0 {
		return "medicamento no creado", http.StatusBadRequest
	}
	createDetailInvoice(invoiceId, IdMedicamentos)
	return "medicamento creado", http.StatusOK
}

func createDetailInvoice(invoiceId int, IdMedicamentos []int) {
	db := database.GetConnection()
	rows := ""
	for index, value := range IdMedicamentos {
		coma := ","
		if index == (len(IdMedicamentos) - 1) {
			coma = ""
		}
		rows += "(" + strconv.Itoa(value) + "," + strconv.Itoa(invoiceId) + ")" + coma
	}
	query := fmt.Sprintf("insert into detalle_factura (id_medicamento, id_factura) values %s", rows)
	db.QueryRow(query)
}
