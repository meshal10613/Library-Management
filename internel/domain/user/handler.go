package user

import (
	"library-management/pkg/httpresponse"

	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetAllUsers(ctx fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Failed to retrieve users",
			Details: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(httpresponse.Success{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}
