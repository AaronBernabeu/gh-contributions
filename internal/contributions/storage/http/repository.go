package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	configuration "github.com/Aaronidas/gh-contributions/internal/configuration"
	contributions "github.com/Aaronidas/gh-contributions/internal/contributions"
)

const (
	url = "https://api.github.com/graphql"
)

type apiContributionRepository struct {
	configurationRepository configuration.ConfigurationRepository
	url                     string
}

func (repo *apiContributionRepository) GetContribution() (*contributions.Contribution, error) {
	token, _ := repo.configurationRepository.GetToken()
	username, _ := repo.configurationRepository.GetUsername()

	payload := payload(username)

	request, _ := http.NewRequest(
		"POST",
		repo.url,
		bytes.NewBuffer(payload),
	)
	request.Header.Set("Authorization", "bearer "+*token)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != 200 {
		fmt.Println("Error response")
		fmt.Println(string(body))
	}

	var data JSONData
	json.Unmarshal(body, &data)

	name := data.GetName()

	if name == "" {
		fmt.Println("Error response")
		fmt.Println(string(body))
	}

	today := data.GetDayContributions()
	week := data.GetWeekContributions()
	month := data.GetMonthContributions()
	year := data.GetTotalContributions()

	return contributions.NewContribution(name, today, week, month, year), error(nil)
}

func NewApiRepository(configurationRepository configuration.ConfigurationRepository) contributions.ContributionRepository {
	return &apiContributionRepository{
		configurationRepository: configurationRepository,
		url:                     url,
	}
}

func payload(username *string) []byte {
	return []byte(`
{
  "query": "query { user(login: \"` + *username + `\") { name contributionsCollection { contributionCalendar { colors totalContributions weeks { contributionDays { color contributionCount date weekday } firstDay } } } }}"
}
`)
}
