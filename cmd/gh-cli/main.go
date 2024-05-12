package main

import (
	viperConfigurationRepository "github.com/Aaronidas/gh-contributions/internal/configuration/viper"
	"github.com/Aaronidas/gh-contributions/internal/contributions/entrypoint"
	apiRepository "github.com/Aaronidas/gh-contributions/internal/contributions/storage/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	initEnv()

	configurationRepository := viperConfigurationRepository.NewViperConfigurationRepository()
	repo := apiRepository.NewApiRepository(configurationRepository)

	rootCmd := &cobra.Command{Use: "gh-cli"}
	rootCmd.AddCommand(entrypoint.InitContributionsCmd(repo))

	rootCmd.Execute()
}

func initEnv() {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err == nil {
		return
	}

	viper.SetConfigFile(".env.dist")
	viper.ReadInConfig()
}
