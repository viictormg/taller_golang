package models

import (
	"fmt"
	"net/http"

	"example.com/taller/src/database"
	"example.com/taller/src/structures"
)

func CreatetMedicine(nombre string, precio float32, ubicacion string) (string, int) {
	db := database.GetConnection()
	var medicineId int
	query := fmt.Sprintf("insert into medicamento(nombre, precio, ubicacion) values ('%s', %f, '%s') returning id", nombre, precio, ubicacion)
	db.QueryRow(query).Scan(&medicineId)

	if medicineId == 0 {
		return "medicamento no creado", http.StatusBadRequest
	}
	return "medicamento creado", http.StatusOK
}

func GetMedicines() ([]structures.Medicine, int) {
	db := database.GetConnection()

	medicines := []structures.Medicine{}
	query := "select id, nombre, precio, ubicacion from medicamento"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
		}
		m := structures.Medicine{}
		if err := rows.Scan(&m.Id, &m.Nombre, &m.Precio, &m.Ubicacion); err != nil {
			panic(err)
		}
		medicines = append(medicines, m)
	}
	return medicines, http.StatusOK
}
