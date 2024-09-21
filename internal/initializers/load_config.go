package initializers

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"

	"github.com/TejasGhatte/go-sail/internal/models"

)
var Config models.Config

func LoadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}

	return nil
}