package structures

type Promotion struct {
	Id          int     `json:"id"`
	Descripcion string  `json:"descripcion"`
	Porcentaje  float32 `json:"porcentaje"`
	FechaInicio string  `json:"fecha_inicio"`
	FechaFin    string  `json:"fecha_fin"`
}
