package goapikey

type (
	// BasicService is the basic service interface for API key validation
	BasicService interface {
		IsAPIKeyValid(apiKey string) bool
	}
)
