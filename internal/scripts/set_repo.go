package scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	RepoURL string `json:"repo_url"`
}

type ConfigManager struct {
	configPath string
}

func NewConfigManager() (*ConfigManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %v", err)
	}

	configDir := filepath.Join(homeDir, ".myapp")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %v", err)
	}

	return &ConfigManager{
		configPath: filepath.Join(configDir, "config.json"),
	}, nil
}

func (cm *ConfigManager) LoadConfig() (*Config, error) {
	config := &Config{}

	data, err := os.ReadFile(cm.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil
		}
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

func (cm *ConfigManager) SaveConfig(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	if err := os.WriteFile(cm.configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}
func RunSetRepo(repoURL string) error {
	cm, err := NewConfigManager()
	if err != nil {
		return fmt.Errorf("failed to initialize config manager: %v", err)
	}

	config, err := cm.LoadConfig()
	if err != nil {
		return err
	}

	// Update repository URL
	config.RepoURL = repoURL

	if err := cm.SaveConfig(config); err != nil {
		return err
	}

	fmt.Printf("Successfully set repository URL to: %s\n", repoURL)
	return nil
}

func GetRepo() (string, error) {
	cm, err := NewConfigManager()
	if err != nil {
		return "", fmt.Errorf("failed to initialize config manager: %v", err)
	}

	config, err := cm.LoadConfig()
	if err != nil {
		return "", err
	}

	return config.RepoURL, nil
}