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
	Create(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type commentHandler struct {
	commentService service.CommentService
	validator      *validator.Validate
}

func NewCommentHandler(commentService service.CommentService, validator *validator.Validate) CommentHandler {
	return &commentHandler{
		commentService: commentService,
		validator:      validator,
	}
}

// @Summary		Create a new comment
// @Description	Create a comment
// @Tags			Comments
// @Accept			json
// @Produce		json
// @Param			comment	body		web.CommentCreate	true	"Comment Data"
// @Success		201		{object}	web.CommentResponse
// @Security		BearerAuth
// @Router			/comments [post]
func (h *commentHandler) Create(ctx echo.Context) error {
	req := new(web.CommentCreate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.commentService.Create(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

// @Summary		Get all comments
// @Description	Retrieve all comments
// @Tags			Comments
// @Accept			json
// @Produce		json
// @Success		200	{array}	web.CommentResponse
// @Security		Bearer
// @Router			/comments [get]
func (h *commentHandler) FindAll(ctx echo.Context) error {
	data, err := h.commentService.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "comments not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Get a comment by ID
// @Description	Retrieve a single comment by ID
// @Tags			Comments
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Comment ID"
// @Success		200	{object}	web.CommentResponse
// @Router			/comments/{id} [get]
func (h *commentHandler) FindByID(ctx echo.Context) error {
	req := new(web.CommentFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.commentService.FindByID(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "comment not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Update an existing comment
// @Description	Update a comment
// @Tags			Comments
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Comment ID"
// @Param			comment	body		web.CommentUpdate	true	"Updated Comment Data"
// @Success		200		{object}	web.CommentResponse
// @Security		BearerAuth
// @Router			/comments/{id} [put]
func (h *commentHandler) Update(ctx echo.Context) error {
	req := new(web.CommentUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.commentService.Update(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "comment not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Delete a comment
// @Description	Remove a comment by ID
// @Tags			Comments
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Comment ID"
// @Success		204
// @Security		BearerAuth
// @Router			/comments/{id} [delete]
func (h *commentHandler) Delete(ctx echo.Context) error {
	req := new(web.CommentDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	req.UserID = uint(claims["user_id"].(float64))

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.commentService.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "comment not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
