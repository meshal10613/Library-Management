package book

import (
	"errors"
	"library-management/internel/domain/book/dto"
	"library-management/pkg/httpresponse"
	"library-management/pkg/querybuilder"
	"library-management/pkg/validation"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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
			Message: "Failed to create book",
			Details: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(httpresponse.Success{
		Success: true,
		Message: "Book created successfully",
		Data:    response,
	})
}

func (h *handler) GetBookByID(ctx fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid book id",
		})
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		if errors.Is(err, ErrBookNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(httpresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Failed to get book",
		})
	}

	return ctx.JSON(httpresponse.Success{
		Success: true,
		Data:    book,
	})
}

func (h *handler) GetAllBooks(ctx fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	books, err := h.service.GetAllBooks(ctx, querybuilder.QueryParams{
		Page:   page,
		Limit:  limit,
		Search: ctx.Query("search"),
		SortBy: ctx.Query("sortBy"),
		Order:  ctx.Query("order"),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Failed to get books",
		})
	}

	if len(books.Data) == 0 {
		return ctx.JSON(httpresponse.Success{
			Success: true,
			Message: "No books found",
		})
	}

	return ctx.JSON(httpresponse.Success{
		Success: true,
		Message: "Books retrived successfully",
		Data:    books.Data,
		Meta:    &books.Meta,
	})
}

func (h *handler) UpdateBook(ctx fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid book id",
		})
	}

	var req dto.UpdateBookRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
		})
	}

	if err := req.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: err.Error(),
		})
	}

	book, err := h.service.UpdateBook(id, &req)
	if err != nil {
		if errors.Is(err, ErrBookNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(httpresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Failed to update book",
		})
	}

	return ctx.JSON(httpresponse.Success{
		Success: true,
		Message: "Book updated successfully",
		Data:    book,
	})
}

func (h *handler) DeleteBook(ctx fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(httpresponse.Error{
			Success: false,
			Message: "Invalid book id",
		})
	}

	err = h.service.DeleteBook(id)
	if err != nil {
		if errors.Is(err, ErrBookNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(httpresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(httpresponse.Error{
			Success: false,
			Message: "Failed to delete book",
		})
	}

	return ctx.JSON(httpresponse.Success{
		Success: true,
		Message: "Book deleted successfully",
	})
}
