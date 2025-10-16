package local

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

type (
	// Service is the local service for API key validation
	Service struct {
		validAPIKeys  map[string]struct{}
		apiKeys       map[string]string
		nameSeparator string
		logger        *slog.Logger
	}
)

// NewService creates a new local service for API key validation
//
// Parameters:
//
//   - logger: the logger to use
//
// Returns:
//
//   - *Service: the local service
//   - error: An error if something went wrong
func NewService(nameSeparator string, logger *slog.Logger) (*Service, error) {
	// Check the name separator
	if nameSeparator == "" {
		return nil, ErrEmptyNameSeparator
	}

	if logger != nil {
		logger = logger.With(
			slog.String("service", "api_key_local"),
		)
	}

	return &Service{
		validAPIKeys:  make(map[string]struct{}),
		apiKeys:       make(map[string]string),
		nameSeparator: nameSeparator,
		logger:        logger,
	}, nil
}

// Load loads the API keys from the file
//
// Parameters:
//
//   - path: the path to the file
//   - logger: the logger to use
//
// Returns:
//
//   - error: An error if something went wrong
func (s Service) Load(path string) error {
	// Open the service API keys file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			if s.logger != nil {
				s.logger.Error(
					"Failed to close service API keys file",
					slog.String("file_path", path),
					slog.String("error", err.Error()),
				)
			}
		}
	}(file)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the line
		line := scanner.Text()

		// Skip empty lines and comments
		if strings.TrimSpace(line) == "" || strings.HasPrefix(
			strings.TrimSpace(line),
			"#",
		) {
			continue
		}

		// Split the line into service name and API key
		parts := strings.SplitN(line, s.nameSeparator, 2)
		if len(parts) != 2 {
			continue // Skip invalid lines
		}

		// Trim spaces
		serviceName := strings.TrimSpace(parts[0])
		apiKey := strings.TrimSpace(parts[1])

		if serviceName == "" || apiKey == "" {
			if s.logger != nil {
				s.logger.Warn(
					"Invalid line in service API keys file",
					slog.String("line", line),
					slog.String("file_path", path),
				)
			}
			continue // Skip invalid lines
		}

		// Add to the map and list
		serviceName = strings.TrimSpace(serviceName)
		s.apiKeys[serviceName] = apiKey
		s.validAPIKeys[apiKey] = struct{}{}
	}
	return scanner.Err()
}

// IsAPIKeyValid checks if the API key is valid
//
// Parameters:
//
//   - apiKey: the API key to check
//
// Returns:
//
// - bool: true if the API key is valid, false otherwise
func (s Service) IsAPIKeyValid(apiKey string) bool {
	_, valid := s.validAPIKeys[apiKey]
	return valid
}
