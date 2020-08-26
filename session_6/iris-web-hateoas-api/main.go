package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/mercadolibre/iris-web-hateoas-api/pkg/contact"
)

func main() {
	app := iris.New()
	app.Use(logger.New())

	svc := contact.NewService()
	h := Handler{cntSvc: svc}
	cntsRoute := app.Party("agenda/contacts")
	cntsRoute.Get("/", h.GetAllContacts)
	cntsRoute.Get("/:cntId", h.GetContact)
	cntsRoute.Post("/", h.CreateContact)
	cntsRoute.Put("/:cntId", h.UpdateContact)
	cntsRoute.Delete("/:cntId", h.DeleteContact)
	cntsRoute.Patch("/:cntId", h.StatusContact)

	app.Run(iris.Addr(":8021"))
}
