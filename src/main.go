package main

import (
	"example.com/taller/src/database"
	"example.com/taller/src/models"
	_ "github.com/lib/pq"
	"github.com/savsgio/atreugo"
)

func main() {
	config := &atreugo.Config{
		Host: "localhost",
		Port: 8000,
	}
	server := atreugo.New(config)

	database.GetConnection()

	// fnMiddlewareOne := func(ctx *atreugo.RequestCtx) (int, error) {
	// 	return fasthttp.StatusOK, nil
	// }

	// fnMiddlewareTwo := func(ctx *atreugo.RequestCtx) (int, error) {
	// 	// Disable this middleware if you don't want to see this error
	// 	return fasthttp.StatusBadRequest, errors.New("Error example")
	// }

	// server.UseMiddleware(fnMiddlewareOne, fnMiddlewareTwo)

	server.Path("GET", "/", func(ctx *atreugo.RequestCtx) error {
		return ctx.HTTPResponse("<h1>Atreugo Micro-Framework</h1>")
	})

	server.Path("GET", "/jsonPage", func(ctx *atreugo.RequestCtx) error {
		return ctx.JSONResponse(atreugo.JSON{"Atreugo": true})
	})
	server.Path("POST", "/createMedicine", func(ctx *atreugo.RequestCtx) error {
		models.Insert("dolex", 2000, "b1")
		response := atreugo.JSON{"status": "ok"}
		return ctx.JSONResponse(response)
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
