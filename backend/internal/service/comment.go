package service

import (
	"net/http"

	"github.com/aternity/zense/internal/entity/domain"
	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/repository"
	"github.com/labstack/echo/v4"
)

type CommentService interface {
	Create(req web.CommentCreate) (*web.CommentResponse, error)
	FindAll() ([]web.CommentResponse, error)
	FindByID(req web.CommentFindByID) (*web.CommentResponse, error)
	Update(req web.CommentUpdate) (*web.CommentResponse, error)
	Delete(req web.CommentDelete) error
}

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) CommentService {
	return &commentService{
		repository: repository,
	}
}

func (s *commentService) Create(req web.CommentCreate) (*web.CommentResponse, error) {
	comment := &domain.Comment{
		UserID:     req.UserID,
		ForumID:    req.ForumID,
		Content:    req.Content,
		Visibility: req.Visibility,
	}

	comment, err := s.repository.Create(comment)
	if err != nil {
		return nil, err
	}

	response := &web.CommentResponse{
		ID:         comment.ID,
		UserID:     comment.UserID,
		ForumID:    comment.ForumID,
		Content:    comment.Content,
		Visibility: comment.Visibility,
		CreatedAt:  &comment.CreatedAt,
	}

	return response, nil
}

func (s *commentService) FindAll() ([]web.CommentResponse, error) {
	comments, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []web.CommentResponse
	for _, comment := range comments {
		responses = append(responses, web.CommentResponse{
			ID:         comment.ID,
			UserID:     comment.UserID,
			ForumID:    comment.ForumID,
			Content:    comment.Content,
			Visibility: comment.Visibility,
			CreatedAt:  &comment.CreatedAt,
			UpdatedAt:  &comment.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *commentService) FindByID(req web.CommentFindByID) (*web.CommentResponse, error) {
	comment, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	response := &web.CommentResponse{
		ID:         comment.ID,
		UserID:     comment.UserID,
		ForumID:    comment.ForumID,
		Content:    comment.Content,
		Visibility: comment.Visibility,
		CreatedAt:  &comment.CreatedAt,
		UpdatedAt:  &comment.UpdatedAt,
	}

	return response, nil
}

func (s *commentService) Update(req web.CommentUpdate) (*web.CommentResponse, error) {
	comment, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	if comment.UserID != req.UserID {
		return nil, echo.NewHTTPError(http.StatusForbidden, "user does not have permission to delete this comment")
	}

	comment = &domain.Comment{
		ID:         req.ID,
		Content:    req.Content,
		Visibility: req.Visibility,
	}

	comment, err = s.repository.Update(comment)
	if err != nil {
		return nil, err
	}

	response := &web.CommentResponse{
		ID:         comment.ID,
		UserID:     comment.UserID,
		ForumID:    comment.ForumID,
		Content:    comment.Content,
		Visibility: comment.Visibility,
		CreatedAt:  &comment.CreatedAt,
		UpdatedAt:  &comment.UpdatedAt,
	}

	return response, nil
}

func (s *commentService) Delete(req web.CommentDelete) error {
	comment, err := s.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if comment.UserID != req.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "user does not have permission to delete this comment")
	}

	comment, err = s.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if err := s.repository.Delete(comment); err != nil {
		return err
	}

	return nil
}
