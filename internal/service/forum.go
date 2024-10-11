package service

import (
	"errors"
	"net/http"

	"github.com/aternity/zense/internal/entity/domain"
	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ForumService interface {
	Create(req web.ForumCreate) (*web.ForumResponse, error)
	FindAll() ([]web.ForumResponse, error)
	FindByID(req web.ForumFindByID) (*web.ForumResponse, error)
	Update(req web.ForumUpdate) (*web.ForumResponse, error)
	Delete(req web.ForumDelete) error
	RemoveTopic(req web.ForumRemoveTopic) error
}

type forumService struct {
	forumRepository repository.ForumRepository
	topicRepository repository.TopicRepository
}

func NewForumService(forumRepository repository.ForumRepository, topicRepository repository.TopicRepository) ForumService {
	return &forumService{
		forumRepository: forumRepository,
		topicRepository: topicRepository,
	}
}

func (s *forumService) Create(req web.ForumCreate) (*web.ForumResponse, error) {
	var topics []domain.Topic
	for _, topicID := range req.Topics {
		topic, err := s.topicRepository.FindByID(topicID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, echo.NewHTTPError(http.StatusNotFound, "topic not found")
			}
			return nil, err
		}
		topics = append(topics, *topic)
	}

	forum := &domain.Forum{
		UserID:  req.UserID,
		Title:   req.Title,
		Topics:  topics,
		Content: req.Content,
	}

	forum, err := s.forumRepository.Create(forum)
	if err != nil {
		return nil, err
	}

	var topicsResponse []web.TopicResponse
	for _, topic := range forum.Topics {
		topicsResponse = append(topicsResponse, web.TopicResponse{
			ID:   topic.ID,
			Name: topic.Name,
		})
	}

	response := &web.ForumResponse{
		ID:        forum.ID,
		UserID:    forum.UserID,
		Title:     forum.Title,
		Topics:    topicsResponse,
		Content:   forum.Content,
		CreatedAt: &forum.CreatedAt,
	}

	return response, nil
}

func (s *forumService) FindAll() ([]web.ForumResponse, error) {
	forums, err := s.forumRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []web.ForumResponse
	for _, forum := range forums {
		var topicsResponse []web.TopicResponse
		for _, topic := range forum.Topics {
			topicsResponse = append(topicsResponse, web.TopicResponse{
				ID:          topic.ID,
				Name:        topic.Name,
				Description: topic.Description,
			})
		}

		responses = append(responses, web.ForumResponse{
			ID:        forum.ID,
			Title:     forum.Title,
			Topics:    topicsResponse,
			Content:   forum.Content,
			CreatedAt: &forum.CreatedAt,
			UpdatedAt: &forum.UpdatedAt,
			User: &web.UserResponse{
				ID:   forum.User.ID,
				Name: forum.User.Name,
			},
		})
	}

	return responses, nil
}

func (s *forumService) FindByID(req web.ForumFindByID) (*web.ForumResponse, error) {
	forum, err := s.forumRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	var topics []web.TopicResponse
	for _, topic := range forum.Topics {
		topics = append(topics, web.TopicResponse{
			ID:          topic.ID,
			Name:        topic.Name,
			Description: topic.Description,
		})
	}

	response := &web.ForumResponse{
		ID:        forum.ID,
		Title:     forum.Title,
		Topics:    topics,
		Content:   forum.Content,
		CreatedAt: &forum.CreatedAt,
		UpdatedAt: &forum.UpdatedAt,
		User: &web.UserResponse{
			ID:   forum.User.ID,
			Name: forum.User.Name,
		},
	}

	return response, nil
}

func (s *forumService) Update(req web.ForumUpdate) (*web.ForumResponse, error) {
	forum, err := s.forumRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	if forum.UserID != req.UserID {
		return nil, echo.NewHTTPError(http.StatusForbidden, "user does not have permission to update this forum")
	}

	var topics []domain.Topic
	for _, topicID := range req.Topics {
		topic, err := s.topicRepository.FindByID(topicID)
		if err != nil {
			return nil, err
		}
		topics = append(topics, *topic)
	}

	forum = &domain.Forum{
		ID:      req.ID,
		Title:   req.Title,
		Topics:  topics,
		Content: req.Content,
	}

	forum, err = s.forumRepository.Update(forum)
	if err != nil {
		return nil, err
	}

	var topicsResponse []web.TopicResponse
	for _, topic := range forum.Topics {
		topicsResponse = append(topicsResponse, web.TopicResponse{
			ID:   topic.ID,
			Name: topic.Name,
		})
	}

	response := &web.ForumResponse{
		ID:        forum.ID,
		UserID:    forum.UserID,
		Title:     forum.Title,
		Topics:    topicsResponse,
		Content:   forum.Content,
		UpdatedAt: &forum.UpdatedAt,
	}

	return response, nil
}

func (s *forumService) Delete(req web.ForumDelete) error {
	forum, err := s.forumRepository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if forum.UserID != req.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "user does not have permission to update this forum")
	}

	if err := s.forumRepository.Delete(forum); err != nil {
		return err
	}

	return nil
}

func (s *forumService) RemoveTopic(req web.ForumRemoveTopic) error {
	forum, err := s.forumRepository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if forum.UserID != req.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "user does not have permission to update this forum")
	}

	if err := s.forumRepository.RemoveTopic(forum); err != nil {
		return err
	}

	return nil
}
