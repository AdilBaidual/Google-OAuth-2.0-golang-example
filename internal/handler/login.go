package handler

import "github.com/gofiber/fiber/v2"

func (h Handler) Login(c *fiber.Ctx) error {
	// create url for auth process.
	// we can pass state as someway to identify
	// and validate the login process.
	URL := h.conf.AuthCodeURL("not-implemented-yet")

	// redirect to the google authentication URL
	return c.Redirect(URL)
}
