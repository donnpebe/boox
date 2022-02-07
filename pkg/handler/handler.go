package handler

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Welcome")
	})
}
