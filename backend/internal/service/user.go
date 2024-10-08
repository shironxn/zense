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
	FindMe(req web.UserFindMe) (*web.UserResponse, error)
	FindAll() ([]web.UserResponse, error)
	FindByID(req web.UserFindByID) (*web.UserResponse, error)
	Update(req web.UserUpdate) (*web.UserResponse, error)
	Delete(req web.UserDelete) error
}

type userService struct {
	repository repository.UserRepository
	jwt        *util.JWT
}

func NewUserService(repository repository.UserRepository, jwt *util.JWT) UserService {
	return &userService{
		repository: repository,
		jwt:        jwt,
	}
}

func (s *userService) Login(req web.UserLogin) (*web.UserAuth, error) {
	user, err := s.repository.FindByEmail(req.Email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}

	token, err := s.jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token")
	}

	response := &web.UserAuth{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return response, nil
}

func (s *userService) Register(req web.UserRegister) (*web.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	user, err = s.repository.Create(user)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
	}

	response := &web.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	return response, nil
}

func (s *userService) FindMe(req web.UserFindMe) (*web.UserResponse, error) {
	user, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	response := &web.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}

	return response, nil
}

func (s *userService) FindAll() ([]web.UserResponse, error) {
	users, err := s.repository.FindAll()
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

func (s *userService) FindByID(req web.UserFindByID) (*web.UserResponse, error) {
	user, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	response := &web.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}

	return response, nil
}

func (s *userService) Update(req web.UserUpdate) (*web.UserResponse, error) {
	user, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	if user.ID != req.UserID {
		return nil, echo.NewHTTPError(http.StatusForbidden, "you do not have permission to update this user")
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
		}
		req.Password = string(hashedPassword)
	}

	user = &domain.User{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err = s.repository.Update(user)
	if err != nil {
		return nil, err
	}

	response := &web.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}

	return response, nil
}

func (s *userService) Delete(req web.UserDelete) error {
	user, err := s.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if user.ID != req.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "you do not have permission to delete this user")
	}

	user, err = s.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if err := s.repository.Delete(user); err != nil {
		return err
	}

	return nil
}
