package service

import (
	"net/http"

	"github.com/aternity/zense/internal/entity/domain"
	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/repository"
	"github.com/labstack/echo/v4"
)

type ForumService interface {
	Create(req web.ForumCreate) (*web.ForumResponse, error)
	FindAll() ([]web.ForumResponse, error)
	FindByID(req web.ForumFindByID) (*web.ForumResponse, error)
	Update(req web.ForumUpdate) (*web.ForumResponse, error)
	Delete(req web.ForumDelete) error
}

type forumService struct {
	repository repository.ForumRepository
}

func NewForumService(repository repository.ForumRepository) ForumService {
	return &forumService{
		repository: repository,
	}
}

func (f *forumService) Create(req web.ForumCreate) (*web.ForumResponse, error) {
	forum := &domain.Forum{
		UserID:  req.UserID,
		Title:   req.Title,
		Topic:   req.Topic,
		Content: req.Content,
	}

	forum, err := f.repository.Create(forum)
	if err != nil {
		return nil, err
	}

	response := &web.ForumResponse{
		ID:        forum.ID,
		UserID:    forum.UserID,
		Title:     forum.Title,
		Topic:     forum.Topic,
		Content:   forum.Content,
		CreatedAt: &forum.CreatedAt,
	}

	return response, nil
}

func (f *forumService) FindAll() ([]web.ForumResponse, error) {
	forums, err := f.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []web.ForumResponse
	for _, forum := range forums {
		responses = append(responses, web.ForumResponse{
			ID:        forum.ID,
			UserID:    forum.UserID,
			Title:     forum.Title,
			Topic:     forum.Topic,
			Content:   forum.Content,
			CreatedAt: &forum.CreatedAt,
			UpdatedAt: &forum.UpdatedAt,
		})
	}

	return responses, nil
}

func (f *forumService) FindByID(req web.ForumFindByID) (*web.ForumResponse, error) {
	forum, err := f.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	response := &web.ForumResponse{
		ID:        forum.ID,
		UserID:    forum.UserID,
		Title:     forum.Title,
		Topic:     forum.Topic,
		Content:   forum.Content,
		CreatedAt: &forum.CreatedAt,
		UpdatedAt: &forum.UpdatedAt,
	}

	return response, nil
}

func (f *forumService) Update(req web.ForumUpdate) (*web.ForumResponse, error) {
	forum, err := f.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	if forum.UserID != req.UserID {
		return nil, echo.NewHTTPError(http.StatusForbidden, "user does not have permission to update this forum")
	}

	forum = &domain.Forum{
		ID:      req.ID,
		Title:   req.Title,
		Topic:   req.Topic,
		Content: req.Content,
	}

	forum, err = f.repository.Update(forum)
	if err != nil {
		return nil, err
	}

	response := &web.ForumResponse{
		ID:        forum.ID,
		UserID:    forum.UserID,
		Title:     forum.Title,
		Topic:     forum.Topic,
		Content:   forum.Content,
		CreatedAt: &forum.CreatedAt,
		UpdatedAt: &forum.UpdatedAt,
	}

	return response, nil
}

func (f *forumService) Delete(req web.ForumDelete) error {
	forum, err := f.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if forum.UserID != req.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "user does not have permission to update this forum")
	}

	forum, err = f.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if err := f.repository.Delete(forum); err != nil {
		return err
	}

	return nil
}
