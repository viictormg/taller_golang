package models

import (
	"fmt"
	"net/http"

	"example.com/taller/src/database"
	"example.com/taller/src/structures"
)

func CreatePromotion(descripcion string, porcentaje float32, fechaInicio string, fechaFin string) (string, int) {
	db := database.GetConnection()
	var promotionId int

	query := fmt.Sprintf("insert into promocion (descripcion, porcentaje, fecha_inicio, fecha_fin) values ('%s', %f, '%s', '%s') returning id", descripcion, porcentaje, fechaInicio, fechaFin)
	if descripcion != "" && porcentaje > 0 && fechaInicio != "" && fechaFin != "" {
		db.QueryRow(query).Scan(&promotionId)
	} else {
		return "promocion no creada, todos los campos son obligatorios", http.StatusBadRequest
	}
	if promotionId == 0 {
		return "promocion no creada", http.StatusBadRequest
	}
	return "promocion creada", http.StatusOK
}

func GetPromotions() ([]structures.Promotion, int) {
	db := database.GetConnection()

	promotions := []structures.Promotion{}
	query := "select id, descripcion, porcentaje, fecha_inicio, fecha_fin from promocion "
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
		}
		p := structures.Promotion{}
		if err := rows.Scan(&p.Id, &p.Descripcion, &p.Porcentaje, &p.FechaInicio, &p.FechaFin); err != nil {
			panic(err)
		}
		promotions = append(promotions, p)
	}
	return promotions, http.StatusOK
}
