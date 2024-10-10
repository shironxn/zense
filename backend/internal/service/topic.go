package service

import (
	"github.com/aternity/zense/internal/entity/domain"
	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/repository"
)

type TopicService interface {
	Create(req web.TopicCreate) (*web.TopicResponse, error)
	FindAll() ([]web.TopicResponse, error)
	FindByID(req web.TopicFindByID) (*web.TopicResponse, error)
	Update(req web.TopicUpdate) (*web.TopicResponse, error)
	Delete(req web.TopicDelete) error
}

type topicService struct {
	repository repository.TopicRepository
}

func NewTopicService(repository repository.TopicRepository) TopicService {
	return &topicService{
		repository: repository,
	}
}

func (s *topicService) Create(req web.TopicCreate) (*web.TopicResponse, error) {
	topic := &domain.Topic{
		Name: req.Name,
	}

	topic, err := s.repository.Create(topic)
	if err != nil {
		return nil, err
	}

	response := &web.TopicResponse{
		ID:   topic.ID,
		Name: topic.Name,
	}

	return response, nil
}

func (s *topicService) FindAll() ([]web.TopicResponse, error) {
	topics, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []web.TopicResponse
	for _, topic := range topics {
		response := web.TopicResponse{
			ID:   topic.ID,
			Name: topic.Name,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *topicService) FindByID(req web.TopicFindByID) (*web.TopicResponse, error) {
	topic, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	response := &web.TopicResponse{
		ID:   topic.ID,
		Name: topic.Name,
	}

	return response, nil
}

func (s *topicService) Update(req web.TopicUpdate) (*web.TopicResponse, error) {
	topic, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	topic.Name = req.Name

	topic, err = s.repository.Update(topic)
	if err != nil {
		return nil, err
	}

	response := &web.TopicResponse{
		ID:   topic.ID,
		Name: topic.Name,
	}

	return response, nil
}

func (s *topicService) Delete(req web.TopicDelete) error {
	topic, err := s.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if err := s.repository.Delete(topic); err != nil {
		return err
	}

	return nil
}
