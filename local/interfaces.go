package local

type (
	// BasicService is the basic service interface for API key validation
	BasicService interface {
		Load(path string) error
		IsAPIKeyValid(apiKey string) bool
	}
)
