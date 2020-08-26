package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/mercadolibre/iris-web-hateoas-api/pkg/contact"
)

type Handler struct {
	cntSvc contact.Service
}

func (h *Handler) GetAllContacts(c context.Context) {
	contacts := h.cntSvc.GetAllContacts()
	for _, cnt := range contacts {
		flush(cnt)
		addLinks(c, cnt)
	}
	c.StatusCode(200)
	c.JSON(contacts)
}

func (h *Handler) GetContact(c context.Context) {
	cntID := c.Params().Get("cntId")
	cnt := h.cntSvc.GetContact(cntID)
	if cnt != nil {
		flush(cnt)
		addLinks(c, cnt)
		c.StatusCode(200)
		c.JSON(cnt)
		return
	}
	c.StatusCode(404)
	c.JSON(iris.Map{"msg": "contact " + cntID + " not found"})
}

func (h *Handler) CreateContact(c context.Context) {

}

func (h *Handler) DeleteContact(c context.Context) {

}

func (h *Handler) UpdateContact(c context.Context) {

}

func (h *Handler) StatusContact(c context.Context) {
	cntID := c.Params().Get("cntId")
	cnt := h.cntSvc.ChangeStatus(cntID)
	if cnt != nil {
		flush(cnt)
		addLinks(c, cnt)
		c.StatusCode(200)
		c.JSON(cnt)
		return
	}
	c.StatusCode(404)
	c.JSON(iris.Map{"msg": "contact " + cntID + " not found"})
}

func addLinks(c context.Context, cnt *contact.Contact) {
	addContactLink(c, cnt)
	addUserLink(c, cnt.UserOwner)
}

func addContactLink(c context.Context, cnt *contact.Contact) {
	basePath := "http://" + c.Host() + "/agenda/contacts"
	self := contact.Link{
		Href: basePath + "/" + cnt.ID,
		Method: "GET",
	}
	links := make(map[string]*contact.Link)
	links["self"] = &self
	if cnt.Status == "active" {
		delete := contact.Link{
			Href: basePath + "/" + cnt.ID,
			Method: "DELETE",
		}
		update := contact.Link{
			Href: basePath + "/" + cnt.ID,
			Method: "PUT",
		}
		disable := contact.Link{
			Href: basePath + "/" + cnt.ID,
			Method: "PATCH",
		}
		links["delete"] = &delete
		links["update"] = &update
		links["disable"] = &disable
	} else {
		enable := contact.Link{
			Href: basePath + "/" + cnt.ID,
			Method: "PATCH",
		}
		links["enable"] = &enable
	}
	cnt.Links = links
}

func addUserLink(c context.Context, user *contact.User) {
	basePath := "http://" + c.Host() + "/users"
	self := contact.Link{
		Href: basePath + "/" + user.ID,
		Method: "GET",
	}
	user.Links = map[string]*contact.Link{"self": &self}
}

func flush(cnt *contact.Contact) {
	cnt.Links = nil
	cnt.UserOwner.Links = nil
}