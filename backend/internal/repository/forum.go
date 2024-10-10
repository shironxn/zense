package repository

import (
	"github.com/aternity/zense/internal/entity/domain"
	"gorm.io/gorm"
)

type ForumRepository interface {
	Create(forum *domain.Forum) (*domain.Forum, error)
	FindAll() ([]domain.Forum, error)
	FindByID(id uint) (*domain.Forum, error)
	Update(forum *domain.Forum) (*domain.Forum, error)
	Delete(forum *domain.Forum) error
	RemoveTopic(forum *domain.Forum) error
}

type forumRepository struct {
	db *gorm.DB
}

func NewForumRepository(db *gorm.DB) ForumRepository {
	return &forumRepository{
		db: db,
	}
}

func (r *forumRepository) Create(forum *domain.Forum) (*domain.Forum, error) {
	if err := r.db.Create(&forum).Error; err != nil {
		return nil, err
	}
	return forum, nil
}

func (r *forumRepository) FindAll() ([]domain.Forum, error) {
	var forums []domain.Forum
	if err := r.db.Preload("User").Preload("Topics").Find(&forums).Error; err != nil {
		return nil, err
	}

	if len(forums) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return forums, nil
}

func (r *forumRepository) FindByID(id uint) (*domain.Forum, error) {
	var forum domain.Forum
	if err := r.db.Preload("User").Preload("Topics").First(&forum, id).Error; err != nil {
		return nil, err
	}
	return &forum, nil
}

func (r *forumRepository) Update(forum *domain.Forum) (*domain.Forum, error) {
	if err := r.db.Updates(&forum).Error; err != nil {
		return nil, err
	}
	return forum, nil
}

func (r *forumRepository) Delete(forum *domain.Forum) error {
	return r.db.Select("Topics").Delete(&forum).Error
}

func (r *forumRepository) RemoveTopic(forum *domain.Forum) error {
	return r.db.Unscoped().Model(&forum).Association("Topics").Unscoped().Clear()
}
