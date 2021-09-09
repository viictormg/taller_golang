package models

import (
	"fmt"
	"net/http"

	"example.com/taller/src/database"
	"github.com/savsgio/atreugo"
)

type Promotion struct {
	Id          int     `json:"id"`
	Descripcion string  `json:"descripcion"`
	Porcentaje  float32 `json:"porcentaje"`
	FechaInicio string  `json:"fecha_inicio"`
	FechaFin    string  `json:"fecha_fin"`
}

func InsertPromotion(descripcion string, porcentaje float32, fechaInicio string, fechaFin string, ctx *atreugo.RequestCtx) error {
	db := database.GetConnection()
	var promotionId int

	query := fmt.Sprintf("insert into promocion (descripcion, porcentaje, fecha_inicio, fecha_fin) values ('%s', %f, '%s', '%s') returning id", descripcion, porcentaje, fechaInicio, fechaFin)
	if descripcion != "" && porcentaje > 0 && fechaInicio != "" && fechaFin != "" {
		db.QueryRow(query).Scan(&promotionId)
	} else {
		return ctx.TextResponse("promocion no creada, todos los campos son obligatorios", http.StatusBadRequest)
	}
	if promotionId == 0 {

		return ctx.TextResponse("promocion no creada", http.StatusBadRequest)

	}
	return ctx.TextResponse("promocion creada", http.StatusOK)
}

func GetPromotions(ctx *atreugo.RequestCtx) error {
	db := database.GetConnection()

	promociones := []Promotion{}
	query := "select id, descripcion, porcentaje, fecha_inicio, fecha_fin from promocion "
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
		}
		p := Promotion{}
		if err := rows.Scan(&p.Id, &p.Descripcion, &p.Porcentaje, &p.FechaInicio, &p.FechaFin); err != nil {
			panic(err)
		}
		promociones = append(promociones, p)
	}
	return ctx.JSONResponse(promociones)

}
