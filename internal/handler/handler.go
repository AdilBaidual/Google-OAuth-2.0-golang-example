package handler

import (
	"Kokos/internal/service"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type Handler struct {
	services *service.Service
	conf     *oauth2.Config
}

func NewHandler(services *service.Service, conf *oauth2.Config) *Handler {
	return &Handler{
		services: services,
		conf:     conf,
	}
}

func (h *Handler) InitRoute() *fiber.App {
	router := fiber.New()

	auth := router.Group("/auth")
	auth.Get("/google", h.Login)
	auth.Get("/callback", h.GoogleCallback)

	api := router.Group("/api")
	api.Get("/test", h.test)
	//api.Post("/withdraw", h.PostWithdrawHandler)
	//api.Post("/transfer", h.PostTransferHandler)
	//api.Get("/balance", h.GetBalanceHandler)

	return router
}

func (h Handler) test(c *fiber.Ctx) error {
	return c.SendString("I'm a GET request!")
}
