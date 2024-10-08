package handler

import (
	"errors"
	"net/http"

	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type userHandler struct {
	service   service.UserService
	validator *validator.Validate
}

func NewUserHandler(service service.UserService, validator *validator.Validate) UserHandler {
	return &userHandler{
		service:   service,
		validator: validator,
	}
}

//	@Summary		User login
//	@Description	Authenticate a user and generate a JWT token
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		web.UserLogin	true	"User Login Request"
//	@Success		200		{object}	web.UserAuth
//	@Router			/auth/login [post]
func (u *userHandler) Login(ctx echo.Context) error {
	req := new(web.UserLogin)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := u.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := u.service.Login(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

//	@Summary		Register a new user
//	@Description	Register a new user with email and password
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		web.UserRegister	true	"User Register Request"
//	@Success		201		{object}	web.UserResponse
//	@Router			/auth/register [post]
func (u *userHandler) Register(ctx echo.Context) error {
	req := new(web.UserRegister)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := u.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := u.service.Register(*req)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return echo.NewHTTPError(http.StatusConflict, "email already registered")
			}
		}

		return err
	}

	return ctx.JSON(http.StatusCreated, data)
}

//	@Summary		Get all users
//	@Description	Retrieve a list of all users
//	@Tags			Users
//	@Produce		json
//	@Success		200	{array}	web.UserResponse
//	@Router			/users [get]
func (u *userHandler) FindAll(ctx echo.Context) error {
	data, err := u.service.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

//	@Summary		Get user by ID
//	@Description	Retrieve a user by their ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	web.UserResponse
//	@Router			/users/{id} [get]
func (u *userHandler) FindByID(ctx echo.Context) error {
	req := new(web.UserFindByID)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := u.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := u.service.FindByID(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

//	@Summary		Update a user
//	@Description	Update a user's information
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"User ID"
//	@Param			user	body		web.UserUpdate	true	"User Update Request"
//	@Success		200		{object}	web.UserResponse
//	@Security		BearerAuth
//	@Router			/users/{id} [put]
func (u *userHandler) Update(ctx echo.Context) error {
	req := new(web.UserUpdate)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	if err := u.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := u.service.Update(*req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}

		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

//	@Summary		Delete a user
//	@Description	Delete a user by their ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Success		204
//	@Security		BearerAuth
//	@Router			/users/{id} [delete]
func (u *userHandler) Delete(ctx echo.Context) error {
	req := new(web.UserDelete)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := u.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := u.service.Delete(*req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}

		return err
	}

	return ctx.NoContent(http.StatusOK)
}
