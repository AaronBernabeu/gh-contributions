package configuration

type ConfigurationRepository interface {
	GetToken() (*string, error)
	GetUsername() (*string, error)
}
