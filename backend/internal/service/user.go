package service

import (
	"net/http"

	"github.com/aternity/zense/internal/entity/domain"
	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/repository"
	"github.com/aternity/zense/internal/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(req web.UserLogin) (*web.UserAuth, error)
	Register(req web.UserRegister) (*web.UserResponse, error)
	FindAll() ([]web.UserResponse, error)
	FindByID(req web.UserFindByID) (*web.UserResponse, error)
	Update(req web.UserUpdate) (*web.UserResponse, error)
	Delete(req web.UserDelete) error
}

type userService struct {
	repository repository.UserRepository
	jwt        util.JWT
}

func NewUserService(repository repository.UserRepository, jwt util.JWT) UserService {
	return &userService{
		repository: repository,
		jwt:        jwt,
	}
}

func (u *userService) Login(req web.UserLogin) (*web.UserAuth, error) {
	user, err := u.repository.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}

	token, err := u.jwt.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &web.UserAuth{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}, nil
}

func (u *userService) Register(req web.UserRegister) (*web.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
	}

	user, err := u.repository.Create(&domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return &web.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (u *userService) FindAll() ([]web.UserResponse, error) {
	users, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []web.UserResponse
	for _, user := range users {
		responses = append(responses, web.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		})
	}

	return responses, nil
}

func (u *userService) FindByID(req web.UserFindByID) (*web.UserResponse, error) {
	user, err := u.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	return &web.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}, nil
}

func (u *userService) Update(req web.UserUpdate) (*web.UserResponse, error) {
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
		}
		req.Password = string(hashedPassword)
	}

	user, err := u.repository.Update(&domain.User{ID: req.ID, Name: req.Name, Email: req.Email, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &web.UserResponse{
		ID:        user.ID,
		UpdatedAt: &user.UpdatedAt,
	}, nil
}

func (u *userService) Delete(req web.UserDelete) error {
	err := u.repository.Delete(req.ID)
	if err != nil {
		return err
	}

	return nil
}
