package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/aternity/zense/internal/entity/web"
	"github.com/google/generative-ai-go/genai"
)

type VentService interface {
	Chat(ctx context.Context, req *web.VentRequest) (*web.VentResponse, error)
	Clear()
}

type ventService struct {
	client              *genai.Client
	conversationHistory []string
}

func NewVentService(client *genai.Client) VentService {
	return &ventService{
		client:              client,
		conversationHistory: []string{},
	}
}

func (s *ventService) Chat(ctx context.Context, req *web.VentRequest) (*web.VentResponse, error) {
	s.conversationHistory = append(s.conversationHistory, fmt.Sprintf("User: %s", req.Message))

	combinedConversation := strings.Join(s.conversationHistory, "\n")

	prompt := fmt.Sprintf(`
    Kamu adalah teman yang dipercaya. Tanggapi pesan berikut dengan empati, gunakan Bahasa Indonesia.
    Berikut adalah percakapan sejauh ini:
    %s
    Pesan terbaru adalah: '%s'. 
    Tolong berikan jawaban yang singkat.
  `, combinedConversation, req.Message)

	model := s.client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	aiResponse := resp.Candidates[0].Content.Parts[0]
	s.conversationHistory = append(s.conversationHistory, fmt.Sprintf("AI: %s", aiResponse))

	response := &web.VentResponse{
		Message: fmt.Sprintf("%s", aiResponse),
	}

	return response, nil
}

func (s *ventService) Clear() {
	s.conversationHistory = []string{}
}
