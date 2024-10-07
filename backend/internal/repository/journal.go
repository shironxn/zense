package repository

import (
	"github.com/aternity/zense/internal/entity/domain"
	"gorm.io/gorm"
)

type JournalRepository interface {
	Create(journal *domain.Journal) (*domain.Journal, error)
	FindAll() ([]domain.Journal, error)
	FindByID(id uint) (*domain.Journal, error)
	Update(journal *domain.Journal) (*domain.Journal, error)
	Delete(journal *domain.Journal) error
}

type journalRepository struct {
	db *gorm.DB
}

func NewJournalRepository(db *gorm.DB) JournalRepository {
	return &journalRepository{
		db: db,
	}
}

func (r *journalRepository) Create(journal *domain.Journal) (*domain.Journal, error) {
	if err := r.db.Create(&journal).Error; err != nil {
		return nil, err
	}
	return journal, nil
}

func (r *journalRepository) FindAll() ([]domain.Journal, error) {
	var journals []domain.Journal
	if err := r.db.Find(&journals).Error; err != nil {
		return nil, err
	}

	if len(journals) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return journals, nil
}

func (r *journalRepository) FindByID(id uint) (*domain.Journal, error) {
	var journal domain.Journal
	if err := r.db.First(&journal, id).Error; err != nil {
		return nil, err
	}
	return &journal, nil
}

func (r *journalRepository) Update(journal *domain.Journal) (*domain.Journal, error) {
	if err := r.db.Updates(&journal).Error; err != nil {
		return nil, err
	}
	return journal, nil
}

func (r *journalRepository) Delete(journal *domain.Journal) error {
	return r.db.Delete(&journal).Error
}
