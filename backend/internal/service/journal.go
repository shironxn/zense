package service

import (
	"net/http"

	"github.com/aternity/zense/internal/entity/domain"
	"github.com/aternity/zense/internal/entity/web"
	"github.com/aternity/zense/internal/repository"
	"github.com/labstack/echo/v4"
)

type JournalService interface {
	Create(req web.JournalCreate) (*web.JournalResponse, error)
	FindAll() ([]web.JournalResponse, error)
	FindByID(req web.JournalFindByID) (*web.JournalResponse, error)
	Update(req web.JournalUpdate) (*web.JournalResponse, error)
	Delete(req web.JournalDelete) error
}

type journalService struct {
	repository repository.JournalRepository
}

func NewJournalService(repository repository.JournalRepository) JournalService {
	return &journalService{
		repository: repository,
	}
}

func (s *journalService) Create(req web.JournalCreate) (*web.JournalResponse, error) {
	journal := &domain.Journal{
		UserID:     req.UserID,
		Mood:       req.Mood,
		Content:    req.Content,
		Visibility: req.Visibility,
	}

	journal, err := s.repository.Create(journal)
	if err != nil {
		return nil, err
	}

	response := &web.JournalResponse{
		ID:         journal.ID,
		UserID:     journal.UserID,
		Mood:       journal.Mood,
		Content:    journal.Content,
		Visibility: journal.Visibility,
		CreatedAt:  &journal.CreatedAt,
	}

	return response, nil
}

func (s *journalService) FindAll() ([]web.JournalResponse, error) {
	journals, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []web.JournalResponse
	for _, journal := range journals {
		response := web.JournalResponse{
			ID:         journal.ID,
			UserID:     journal.UserID,
			Mood:       journal.Mood,
			Content:    journal.Content,
			Visibility: journal.Visibility,
			CreatedAt:  &journal.CreatedAt,
			UpdatedAt:  &journal.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *journalService) FindByID(req web.JournalFindByID) (*web.JournalResponse, error) {
	journal, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	response := &web.JournalResponse{
		ID:         journal.ID,
		UserID:     journal.UserID,
		Mood:       journal.Mood,
		Content:    journal.Content,
		Visibility: journal.Visibility,
		CreatedAt:  &journal.CreatedAt,
		UpdatedAt:  &journal.UpdatedAt,
	}

	return response, nil
}

func (s *journalService) Update(req web.JournalUpdate) (*web.JournalResponse, error) {
	journal, err := s.repository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	if journal.UserID != req.UserID {
		return nil, echo.NewHTTPError(http.StatusForbidden, "user does not have permission to update this journal")
	}

	journal = &domain.Journal{
		ID:         req.ID,
		UserID:     req.UserID,
		Mood:       req.Mood,
		Content:    req.Content,
		Visibility: req.Visibility,
	}

	journal, err = s.repository.Update(journal)
	if err != nil {
		return nil, err
	}

	response := &web.JournalResponse{
		ID:         journal.ID,
		UserID:     journal.UserID,
		Mood:       journal.Mood,
		Content:    journal.Content,
		Visibility: journal.Visibility,
		UpdatedAt:  &journal.UpdatedAt,
	}

	return response, nil
}

func (s *journalService) Delete(req web.JournalDelete) error {
	journal, err := s.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if journal.UserID != req.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "user does not have permission to delete this journal")
	}

	journal, err = s.repository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if err := s.repository.Delete(journal); err != nil {
		return err
	}

	return nil
}
