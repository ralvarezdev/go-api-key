package local

import (
	goapikey "github.com/ralvarezdev/go-api-key"
)

type (
	// BasicService is the basic service interface for API key validation
	BasicService interface {
		goapikey.BasicService
		Load(path string) error
	}
)
