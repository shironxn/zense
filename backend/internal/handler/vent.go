package handler

import (
	"context"
	"net/http"

	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type VentHandler interface {
	Chat(ctx echo.Context) error
	Clear(ctx echo.Context) error
}

type ventHandler struct {
	ventService service.VentService
	validator   *validator.Validate
}

func NewVentHandler(ventService service.VentService, validator *validator.Validate) VentHandler {
	return &ventHandler{
		ventService: ventService,
		validator:   validator,
	}
}

// @Summary		Chat with AI
// @Description	Start a chat session with AI for venting (confiding)
// @Tags			Vents
// @Accept			json
// @Produce		json
// @Param			chat	body		web.VentRequest	true	"Vent Request"
// @Success		200		{object}	web.VentResponse
// @Security		BearerAuth
// @Router			/vents [post]
func (h *ventHandler) Chat(ctx echo.Context) error {
	req := new(web.VentRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c := context.Background()
	data, err := h.ventService.Chat(c, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

// @Summary		Clear chat history
// @Description	Clear the chat history for the current user
// @Tags			Vents
// @Success		204
// @Security		BearerAuth
// @Router			/vents [delete]
func (h *ventHandler) Clear(ctx echo.Context) error {
	h.ventService.Clear()
	return ctx.NoContent(http.StatusNoContent)
}
