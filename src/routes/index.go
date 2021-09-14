package routes

import (
	"github.com/savsgio/atreugo/v11"
)

func SetRouter(router *atreugo.Router) {
	router.GET("/medicamento", GetMedicine)
	router.POST("/medicamento", CreateMedicine)
	//
	router.GET("/promocion", GetPromotions)
	router.POST("/promocion", CreatePromotion)
	//
	router.GET("/factura", GetInvoice)
	router.POST("/factura", CreateInvoice)
	router.GET("/factura/simular", SimulateInvoice)

}
