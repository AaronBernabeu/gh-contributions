package configuration

import (
	"os"

	configuration "github.com/Aaronidas/gh-contributions/internal/configuration"
	viper "github.com/spf13/viper"
)

const (
	tokenKey    = "GH_TOKEN"
	usernameKey = "GH_USERNAME"
)

type viperConfigurationRepository struct {
}

func (repo *viperConfigurationRepository) GetToken() (*string, error) {
	token := os.Getenv(tokenKey)

	if token == "" {
		token = viper.GetString(tokenKey)
	}

	return &token, nil
}

func (repo *viperConfigurationRepository) GetUsername() (*string, error) {
	username := os.Getenv(usernameKey)

	if username == "" {
		username = viper.GetString(usernameKey)
	}

	return &username, nil
}

func NewViperConfigurationRepository() configuration.ConfigurationRepository {
	return &viperConfigurationRepository{}
}
