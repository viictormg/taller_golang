package models

import (
	"fmt"
	"net/http"

	"example.com/taller/src/database"
	"github.com/savsgio/atreugo"
)

type Medicine struct {
	Id        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	Precio    float32 `json:"precio"`
	Ubicacion string  `json:"ubicacion"`
}

func InsertMedicine(nombre string, precio float32, ubicacion string, ctx *atreugo.RequestCtx) error {
	db := database.GetConnection()
	var medicineId int

	query := fmt.Sprintf("insert into medicamento(nombre, precio, ubicacion) values ('%s', %f, '%s') returning id", nombre, precio, ubicacion)
	db.QueryRow(query).Scan(&medicineId)

	if medicineId == 0 {
		return ctx.TextResponse("medicamento no creado", http.StatusBadRequest)
	}
	return ctx.TextResponse("medicamento creado", http.StatusOK)
}

func GetMedicines(ctx *atreugo.RequestCtx) error {
	db := database.GetConnection()

	medicines := []Medicine{}
	query := "select id, nombre, precio, ubicacion from medicamento"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		if rows.Err() != nil {
			// panic(rows.Err())
			fmt.Println(rows.Err())
		}
		m := Medicine{}
		if err := rows.Scan(&m.Id, &m.Nombre, &m.Precio, &m.Ubicacion); err != nil {
			panic(err)
		}
		medicines = append(medicines, m)
	}
	return ctx.JSONResponse(medicines)
}
