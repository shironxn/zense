package handler

import (
	"errors"
	"net/http"

	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TopicHandler interface {
	Create(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	FindByID(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type topicHandler struct {
	topicService service.TopicService
	validator    *validator.Validate
}

func NewTopicHandler(topicService service.TopicService, validator *validator.Validate) TopicHandler {
	return &topicHandler{
		topicService: topicService,
		validator:    validator,
	}
}

// @Summary		Create Topic
// @Description	Create a new topic
// @Tags			Topics
// @Accept			json
// @Produce		json
// @Param			topic	body		web.TopicCreate	true	"Topic Data"
// @Success		201		{object}	web.TopicResponse
// @Security		BearerAuth
// @Router			/topics [post]
func (h *topicHandler) Create(ctx echo.Context) error {
	req := new(web.TopicCreate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.topicService.Create(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

// @Summary		Get All Topics
// @Description	Get all topics
// @Tags			Topics
// @Produce		json
// @Success		200	{array}	web.TopicResponse
// @Router			/topics [get]
func (h *topicHandler) FindAll(ctx echo.Context) error {
	data, err := h.topicService.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "topics not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Get Topic by ID
// @Description	Get topic by its ID
// @Tags			Topics
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Topic ID"
// @Success		200	{object}	web.TopicResponse
// @Router			/topics/{id} [get]
func (h *topicHandler) FindByID(ctx echo.Context) error {
	req := new(web.TopicFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.topicService.FindByID(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "topic not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Update Topic
// @Description	Update an existing topic
// @Tags			Topics
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"Topic ID"
// @Param			topic	body		web.TopicUpdate	true	"Updated Topic Data"
// @Success		200		{object}	web.TopicResponse
// @Security		BearerAuth
// @Router			/topics/{id} [put]
func (h *topicHandler) Update(ctx echo.Context) error {
	req := new(web.TopicUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.topicService.Update(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "topic not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Delete Topic
// @Description	Delete a topic
// @Tags			Topics
// @Accept			json
// @Param			id	path	int	true	"Topic ID"
// @Success		204
// @Security		BearerAuth
// @Router			/topics/{id} [delete]
func (h *topicHandler) Delete(ctx echo.Context) error {
	req := new(web.TopicDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.topicService.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "topic not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
