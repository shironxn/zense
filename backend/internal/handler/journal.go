package handler

import (
	"errors"
	"net/http"

	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type JournalHandler interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type journalHandler struct {
	journalService service.JournalService
	validator      *validator.Validate
}

func NewJournalHandler(journalService service.JournalService, validator *validator.Validate) JournalHandler {
	return &journalHandler{
		journalService: journalService,
		validator:      validator,
	}
}

// @Summary		Create Journal
// @Description	Create a new journal entry
// @Tags			Journals
// @Accept			json
// @Produce		json
// @Param			journal	body		web.JournalCreate	true	"Journal Data"
// @Success		201		{object}	web.JournalResponse
// @Security		BearerAuth
// @Router			/journals [post]
func (h *journalHandler) Create(ctx echo.Context) error {
	req := new(web.JournalCreate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.journalService.Create(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

// @Summary		Get All Journals
// @Description	Get all journals
// @Tags			Journals
// @Produce		json
// @Success		200	{array}	web.JournalResponse
// @Router			/journals [get]
func (h *journalHandler) FindAll(ctx echo.Context) error {
	data, err := h.journalService.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "journals not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Get Journal by ID
// @Description	Get journal by its ID
// @Tags			Journals
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Journal ID"
// @Success		200	{object}	web.JournalResponse
// @Router			/journals/{id} [get]
func (h *journalHandler) FindByID(ctx echo.Context) error {
	req := new(web.JournalFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.journalService.FindByID(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "journal not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Update Journal
// @Description	Update an existing journal entry
// @Tags			Journals
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Journal ID"
// @Param			journal	body		web.JournalUpdate	true	"Updated Journal Data"
// @Success		200		{object}	web.JournalResponse
// @Security		BearerAuth
// @Router			/journals/{id} [put]
func (h *journalHandler) Update(ctx echo.Context) error {
	req := new(web.JournalUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.journalService.Update(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "journal not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Delete Journal
// @Description	Delete a journal entry
// @Tags			Journals
// @Accept			json
// @Param			id	path	int	true	"Journal ID"
// @Success		204
// @Security		BearerAuth
// @Router			/journals/{id} [delete]
func (h *journalHandler) Delete(ctx echo.Context) error {
	req := new(web.JournalDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.journalService.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "journal not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusOK)
}
