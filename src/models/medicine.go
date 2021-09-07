package models

import (
	"fmt"

	"example.com/taller/src/database"
)

type Medicine struct {
	id        int     `json:id`
	nombre    string  `json:nombre`
	precio    float32 `json:precio`
	ubicacion string  `json:ubicacion`
}

func Insert(nombre string, precio float32, ubicacion string) (Medicine, bool) {
	db := database.GetConnection()
	var medicineId int

	query := fmt.Sprintf("insert into medicamento(nombre, precio, ubicacion) values ('%s', %f, '%s') returning id", nombre, precio, ubicacion)
	db.QueryRow(query).Scan(&medicineId)

	if medicineId == 0 {
		return Medicine{}, false
	}
	return Medicine{medicineId, nombre, precio, ubicacion}, true
}
