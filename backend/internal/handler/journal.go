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
	service   service.JournalService
	validator *validator.Validate
}

func NewJournalHandler(service service.JournalService, validator *validator.Validate) JournalHandler {
	return &journalHandler{
		service:   service,
		validator: validator,
	}
}

// @Summary		Create Journal
// @Description	Create a new journal entry
// @Tags			Journal
// @Accept			json
// @Produce		json
// @Param			journal	body		web.JournalCreate	true	"Journal Data"
// @Success		201		{object}	web.JournalResponse
// @Security		BearerAuth
// @Router			/journals [post]
func (j *journalHandler) Create(ctx echo.Context) error {
	req := new(web.JournalCreate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := j.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := j.service.Create(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

// @Summary		Get All Journals
// @Description	Get all journals
// @Tags			Journal
// @Produce		json
// @Success		200	{array}	web.JournalResponse
// @Router			/journals [get]
func (j *journalHandler) FindAll(ctx echo.Context) error {
	data, err := j.service.FindAll()
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
// @Tags			Journal
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Journal ID"
// @Success		200	{object}	web.JournalResponse
// @Router			/journals/{id} [get]
func (j *journalHandler) FindByID(ctx echo.Context) error {
	req := new(web.JournalFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := j.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := j.service.FindByID(*req)
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
// @Tags			Journal
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Journal ID"
// @Param			journal	body		web.JournalUpdate	true	"Updated Journal Data"
// @Success		200		{object}	web.JournalResponse
// @Security		BearerAuth
// @Router			/journals/{id} [put]
func (j *journalHandler) Update(ctx echo.Context) error {
	req := new(web.JournalUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := j.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := j.service.Update(*req)
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
// @Tags			Journal
// @Accept			json
// @Param			id	path	int	true	"Journal ID"
// @Success		204
// @Security		BearerAuth
// @Router			/journals/{id} [delete]
func (j *journalHandler) Delete(ctx echo.Context) error {
	req := new(web.JournalDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := j.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := j.service.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "journal not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusOK)
}
