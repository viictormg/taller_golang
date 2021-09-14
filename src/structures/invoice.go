package structures

type Invoice struct {
	Id             int     `json:"id"`
	FechaCreacion  string  `json:"fecha_creacion"`
	PagoTotal      float32 `json:"pago_total"`
	IdPromocion    int     `json:"id_promocion"`
	IdMedicamentos []int   `json:"id_medicamentos"`
}

type InvoiceResponseGet struct {
	Id             int
	FechaCreacion  string
	PagoTotal      float32
	Promocion      Promotion
	IdMedicamentos []Medicine
}
