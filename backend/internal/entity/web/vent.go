package web

type VentResponse struct {
	Message string `json:"message"`
}

type VentRequest struct {
	Message string `json:"message" validate:"required"`
}
