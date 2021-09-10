package structures

type Invoice struct {
	Id             int     `json:"id"`
	FechaCreacion  string  `json:"fecha_creacion"`
	PagoTotal      float32 `json:"pago_total"`
	IdPromocion    int     `json:"id_promocion"`
	IdMedicamentos []int   `json:"id_medicamentos"`
}
