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
	forumService service.ForumService
	validator    *validator.Validate
}

func NewForumHandler(forumService service.ForumService, validator *validator.Validate) ForumHandler {
	return &forumHandler{
		forumService: forumService,
		validator:    validator,
	}
}

// @Summary		Create Forum
// @Description	Create a new forum post
// @Tags			Forums
// @Accept			json
// @Produce		json
// @Param			forum	body		web.ForumCreate	true	"Forum Data"
// @Success		201		{object}	web.ForumResponse
// @Security		BearerAuth
// @Router			/forums [post]
func (h *forumHandler) Create(ctx echo.Context) error {
	req := new(web.ForumCreate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.forumService.Create(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

// @Summary		Get All Forums
// @Description	Get all forum posts
// @Tags			Forums
// @Produce		json
// @Success		200	{array}	web.ForumResponse
// @Router			/forums [get]
func (h *forumHandler) FindAll(ctx echo.Context) error {
	data, err := h.forumService.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "forums not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Get Forum by ID
// @Description	Get a forum post by ID
// @Tags			Forums
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Forum ID"
// @Success		200	{object}	web.ForumResponse
// @Router			/forums/{id} [get]
func (h *forumHandler) FindByID(ctx echo.Context) error {
	req := new(web.ForumFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.forumService.FindByID(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "forum not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Update Forum
// @Description	Update a forum post
// @Tags			Forums
// @Accept			json
// @Produce		json
// @Param			id		path		string			true	"Forum ID"
// @Param			forum	body		web.ForumUpdate	true	"Updated Forum Data"
// @Success		200		{object}	web.ForumResponse
// @Security		BearerAuth
// @Router			/forums/{id} [put]
func (h *forumHandler) Update(ctx echo.Context) error {
	req := new(web.ForumUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.forumService.Update(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "forum not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Delete Forum
// @Description	Delete a forum post
// @Tags			Forums
// @Accept			json
// @Param			id	path	string	true	"Forum ID"
// @Success		204
// @Security		BearerAuth
// @Router			/forums/{id} [delete]
func (h *forumHandler) Delete(ctx echo.Context) error {
	req := new(web.ForumDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.forumService.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "forum not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusOK)
}
