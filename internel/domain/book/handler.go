package book

import (
	"errors"
	"library-management/internel/domain/book/dto"
	"library-management/pkg/httpresponse"
	"library-management/pkg/validation"

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

func (h *handler) CreateBook(ctx fiber.Ctx) error {
	var req dto.CreateBookRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	validator, ok := ctx.Locals("validator").(*validation.CustomValidator)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Validator not available",
		})
	}

	if err := validator.Validate(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.CreateBook(&req)
	if err != nil {
		if errors.Is(err, ErrorAlreadyExist) {
			return ctx.Status(fiber.StatusConflict).JSON(httpresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Failed to register user",
			Details: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "Book created successfully",
		Data:    response,
	})
}
