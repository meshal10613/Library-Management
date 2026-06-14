package auth

import (
	"errors"
	"fmt"
	"library-management/internel/domain/auth/dto"
	"library-management/pkg/httpresponse"
	"library-management/pkg/utils"
	"library-management/pkg/validation"

	"github.com/go-playground/validator/v10"
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

func formatValidationErrors(err error) string {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return err.Error()
	}

	msg := ""
	for i, fe := range ve {
		if i > 0 {
			msg += "; "
		}
		switch fe.Tag() {
		case "required":
			msg += fmt.Sprintf("'%s' is required", fe.Field())
		case "email":
			msg += fmt.Sprintf("'%s' must be a valid email address", fe.Field())
		case "min":
			msg += fmt.Sprintf("'%s' must be at least %s characters", fe.Field(), fe.Param())
		case "max":
			msg += fmt.Sprintf("'%s' must be at most %s characters", fe.Field(), fe.Param())
		case "password":
			msg += fmt.Sprintf("'%s' must be at least 8 characters and contain uppercase, lowercase, number and special character", fe.Field())
		default:
			msg += fmt.Sprintf("'%s' failed validation (%s)", fe.Field(), fe.Tag())
		}
	}
	return msg
}

func (h *handler) RegisterUser(ctx fiber.Ctx) error {
	var req dto.RegisterUserRequest
 
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}
 
	cv, ok := ctx.Locals("validator").(*validation.CustomValidator)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Validator not available",
		})
	}
 
	if err := cv.Validate(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Validation failed",
			Details: formatValidationErrors(err),
		})
	}
 
	response, err := h.service.RegisterUser(&req)
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
 
	utils.SetAuthCookie(ctx, response.Token)
 
	return ctx.Status(fiber.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "User registered successfully",
		Data:    response,
	})
}
