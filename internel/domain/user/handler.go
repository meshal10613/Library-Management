package user

import (
	"library-management/pkg/httpresponse"
	"library-management/pkg/querybuilder"
	"strconv"

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
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	result, err := h.service.GetAllUsers(ctx, querybuilder.QueryParams{
		Page:   page,
		Limit:  limit,
		Search: ctx.Query("search"),
		SortBy: ctx.Query("sortBy"),
		Order:  ctx.Query("order"),
	})
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
		Data:    result.Data,
		Meta: &result.Meta,
	})
}
