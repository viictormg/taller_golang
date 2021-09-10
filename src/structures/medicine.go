package structures

type Medicine struct {
	Id        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	Precio    float32 `json:"precio"`
	Ubicacion string  `json:"ubicacion"`
}
