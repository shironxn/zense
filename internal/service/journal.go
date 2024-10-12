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
	journalRepository repository.JournalRepository
}

func NewJournalService(journalRepository repository.JournalRepository) JournalService {
	return &journalService{
		journalRepository: journalRepository,
	}
}

func (s *journalService) Create(req web.JournalCreate) (*web.JournalResponse, error) {
	journal := &domain.Journal{
		UserID:     req.UserID,
		Mood:       req.Mood,
		Content:    req.Content,
		Visibility: req.Visibility,
	}

	journal, err := s.journalRepository.Create(journal)
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
	journals, err := s.journalRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []web.JournalResponse
	for _, journal := range journals {
		response := web.JournalResponse{
			ID:         journal.ID,
			Mood:       journal.Mood,
			Content:    journal.Content,
			Visibility: journal.Visibility,
			CreatedAt:  &journal.CreatedAt,
			UpdatedAt:  &journal.UpdatedAt,
			User: &web.UserResponse{
				ID:   journal.User.ID,
				Name: journal.User.Name,
			},
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *journalService) FindByID(req web.JournalFindByID) (*web.JournalResponse, error) {
	journal, err := s.journalRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	response := &web.JournalResponse{
		ID:         journal.ID,
		Mood:       journal.Mood,
		Content:    journal.Content,
		Visibility: journal.Visibility,
		CreatedAt:  &journal.CreatedAt,
		UpdatedAt:  &journal.UpdatedAt,
		User: &web.UserResponse{
			ID:   journal.User.ID,
			Name: journal.User.Name,
		},
	}

	return response, nil
}

func (s *journalService) Update(req web.JournalUpdate) (*web.JournalResponse, error) {
	journal, err := s.journalRepository.FindByID(req.ID)
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

	journal, err = s.journalRepository.Update(journal)
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
	journal, err := s.journalRepository.FindByID(req.ID)
	if err != nil {
		return err
	}

	if journal.UserID != req.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "user does not have permission to delete this journal")
	}

	if err := s.journalRepository.Delete(journal); err != nil {
		return err
	}

	return nil
}
