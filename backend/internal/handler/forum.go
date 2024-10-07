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

type ForumHandler interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type forumHandler struct {
	service   service.ForumService
	validator *validator.Validate
}

func NewForumHandler(service service.ForumService, validator *validator.Validate) ForumHandler {
	return &forumHandler{
		service:   service,
		validator: validator,
	}
}

func (f *forumHandler) Create(ctx echo.Context) error {
	req := new(web.ForumCreate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := f.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := f.service.Create(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

func (f *forumHandler) FindAll(ctx echo.Context) error {
	data, err := f.service.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "forums not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (f *forumHandler) FindByID(ctx echo.Context) error {
	req := new(web.ForumFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := f.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := f.service.FindByID(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "forum not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (f *forumHandler) Update(ctx echo.Context) error {
	req := new(web.ForumUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := f.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := f.service.Update(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "forum not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (f *forumHandler) Delete(ctx echo.Context) error {
	req := new(web.ForumDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := f.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := f.service.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "forum not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusOK)
}
