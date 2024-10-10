package repository

import (
	"github.com/aternity/zense/internal/entity/domain"
	"gorm.io/gorm"
)

type TopicRepository interface {
	Create(topic *domain.Topic) (*domain.Topic, error)
	FindAll() ([]domain.Topic, error)
	FindByID(id uint) (*domain.Topic, error)
	Update(topic *domain.Topic) (*domain.Topic, error)
	Delete(topic *domain.Topic) error
}

type topicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) TopicRepository {
	return &topicRepository{
		db: db,
	}
}

func (r *topicRepository) Create(topic *domain.Topic) (*domain.Topic, error) {
	if err := r.db.Create(&topic).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *topicRepository) FindAll() ([]domain.Topic, error) {
	var topics []domain.Topic
	if err := r.db.Find(&topics).Error; err != nil {
		return nil, err
	}

	if len(topics) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return topics, nil
}

func (r *topicRepository) FindByID(id uint) (*domain.Topic, error) {
	var topic domain.Topic
	if err := r.db.First(&topic, id).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *topicRepository) Update(topic *domain.Topic) (*domain.Topic, error) {
	if err := r.db.Updates(&topic).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *topicRepository) Delete(topic *domain.Topic) error {
	return r.db.Delete(&topic).Error
}
