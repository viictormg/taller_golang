package models

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/taller/src/database"
	"example.com/taller/src/structures"
)

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

func GetInvoice(startDate string, endDate string) ([]structures.InvoiceResponseGet, int) {
	db := database.GetConnection()

	invoices := []structures.InvoiceResponseGet{}
	const query = "select tf.id, tf.fecha_creacion, tf.pago_total, tf.id_promocion from factura tf"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
		}
		invoice := structures.InvoiceResponseGet{}
		var promotionId int

		if err := rows.Scan(&invoice.Id, &invoice.FechaCreacion, &invoice.PagoTotal, &promotionId); err != nil {
			panic(err)
		}
		invoice.Promocion = GetPromotionById(promotionId)
		invoice.IdMedicamentos = GetMedicinesByInvoice(invoice.Id)
		invoices = append(invoices, invoice)
	}
	return invoices, http.StatusOK
}

func SimulateInvoice(date string, IdMedicamentos []string) (float32, int) {
	descuento := GetDiscountByDate(date)
	var porcentajeDescuento float32
	var totalConDescuento float32
	var total int
	for _, id := range IdMedicamentos {
		valor := GetMedicineById(id)
		total += int(valor.Precio)
	}

	porcentajeDescuento = (float32(descuento) / 100)
	fmt.Println("porcentajeDescuento", porcentajeDescuento)

	totalConDescuento = (float32(total) - (float32(total) * porcentajeDescuento))
	return totalConDescuento, http.StatusOK
}
