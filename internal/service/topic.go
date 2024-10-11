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
	topicRepository repository.TopicRepository
}

func NewTopicService(topicRepository repository.TopicRepository) TopicService {
	return &topicService{
		topicRepository: topicRepository,
	}
}

func (s *topicService) Create(req web.TopicCreate) (*web.TopicResponse, error) {
	topic := &domain.Topic{
		Name: req.Name,
    Description: req.Description,
	}

	topic, err := s.topicRepository.Create(topic)
	if err != nil {
		return nil, err
	}

	response := &web.TopicResponse{
		ID:   topic.ID,
		Name: topic.Name,
    Description: topic.Description,
    CreatedAt: &topic.CreatedAt,
	}

	return response, nil
}

func (s *topicService) FindAll() ([]web.TopicResponse, error) {
	topics, err := s.topicRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []web.TopicResponse
	for _, topic := range topics {
		response := web.TopicResponse{
			ID:   topic.ID,
			Name: topic.Name,
      Description: topic.Description,
      CreatedAt: &topic.CreatedAt,
      UpdatedAt: &topic.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *topicService) FindByID(req web.TopicFindByID) (*web.TopicResponse, error) {
	topic, err := s.topicRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	response := &web.TopicResponse{
		ID:   topic.ID,
		Name: topic.Name,
    Description: topic.Description,
    CreatedAt: &topic.CreatedAt,
    UpdatedAt: &topic.UpdatedAt,
	}

	return response, nil
}

func (s *topicService) Update(req web.TopicUpdate) (*web.TopicResponse, error) {
	topic, err := s.topicRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

  topic = &domain.Topic{
    ID: req.ID,
    Name: req.Name,
    Description: req.Description,
  }

	topic, err = s.topicRepository.Update(topic)
	if err != nil {
		return nil, err
	}

	response := &web.TopicResponse{
		ID:   topic.ID,
		Name: topic.Name,
    Description: topic.Description,
    UpdatedAt: &topic.UpdatedAt,
	}

	return response, nil
}

func (s *topicService) Delete(req web.TopicDelete) error {
	topic, err := s.topicRepository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if err := s.topicRepository.Delete(topic); err != nil {
		return err
	}

	return nil
}

