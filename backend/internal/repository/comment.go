package repository

import (
	"github.com/aternity/zense/internal/entity/domain"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *domain.Comment) (*domain.Comment, error)
	FindAll() ([]domain.Comment, error)
	FindByID(id uint) (*domain.Comment, error)
	Update(comment *domain.Comment) (*domain.Comment, error)
	Delete(comment *domain.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(comment *domain.Comment) (*domain.Comment, error) {
	if err := r.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentRepository) FindAll() ([]domain.Comment, error) {
	var comments []domain.Comment
	if err := r.db.Find(&comments).Error; err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return comments, nil
}

func (r *commentRepository) FindByID(id uint) (*domain.Comment, error) {
	var comment domain.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) Update(comment *domain.Comment) (*domain.Comment, error) {
	if err := r.db.Updates(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentRepository) Delete(comment *domain.Comment) error {
	return r.db.Delete(&comment).Error
}
