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
