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

type CommentHandler interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type commentHandler struct {
	service   service.CommentService
	validator *validator.Validate
}

func NewCommentHandler(service service.CommentService, validator *validator.Validate) CommentHandler {
	return &commentHandler{
		service:   service,
		validator: validator,
	}
}

func (c *commentHandler) Create(ctx echo.Context) error {
	req := new(web.CommentCreate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := c.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := c.service.Create(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

func (c *commentHandler) FindAll(ctx echo.Context) error {
	data, err := c.service.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "comments not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c *commentHandler) FindByID(ctx echo.Context) error {
	req := new(web.CommentFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := c.service.FindByID(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "comment not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c *commentHandler) Update(ctx echo.Context) error {
	req := new(web.CommentUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := c.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := c.service.Update(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "comment not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c *commentHandler) Delete(ctx echo.Context) error {
	req := new(web.CommentDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := c.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.service.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "comment not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusOK)
}
