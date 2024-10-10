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
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type topicHandler struct {
	service   service.TopicService
	validator *validator.Validate
}

func NewTopicHandler(service service.TopicService, validator *validator.Validate) TopicHandler {
	return &topicHandler{
		service:   service,
		validator: validator,
	}
}

// @Summary		Create Topic
// @Description	Create a new topic
// @Tags			Topic
// @Accept			json
// @Produce		json
// @Param			topic	body		web.TopicCreate	true	"Topic Data"
// @Success		201		{object}	web.TopicResponse
// @Router			/topics [post]
func (t *topicHandler) Create(ctx echo.Context) error {
	req := new(web.TopicCreate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := t.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := t.service.Create(*req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

// @Summary		Get All Topics
// @Description	Get all topics
// @Tags			Topic
// @Produce		json
// @Success		200	{array}	web.TopicResponse
// @Router			/topics [get]
func (t *topicHandler) FindAll(ctx echo.Context) error {
	data, err := t.service.FindAll()
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
// @Tags			Topic
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Topic ID"
// @Success		200	{object}	web.TopicResponse
// @Router			/topics/{id} [get]
func (t *topicHandler) FindByID(ctx echo.Context) error {
	req := new(web.TopicFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := t.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := t.service.FindByID(*req)
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
// @Tags			Topic
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"Topic ID"
// @Param			topic	body		web.TopicUpdate	true	"Updated Topic Data"
// @Success		200		{object}	web.TopicResponse
// @Router			/topics/{id} [put]
func (t *topicHandler) Update(ctx echo.Context) error {
	req := new(web.TopicUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := t.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := t.service.Update(*req)
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
// @Tags			Topic
// @Accept			json
// @Param			id	path	int	true	"Topic ID"
// @Success		204
// @Router			/topics/{id} [delete]
func (t *topicHandler) Delete(ctx echo.Context) error {
	req := new(web.TopicDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := t.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := t.service.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "topic not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusOK)
}
