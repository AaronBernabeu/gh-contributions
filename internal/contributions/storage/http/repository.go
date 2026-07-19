package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	configuration "github.com/AaronBernabeu/gh-contributions/internal/configuration"
	contributions "github.com/AaronBernabeu/gh-contributions/internal/contributions"
)

const (
	url = "https://api.github.com/graphql"

	// recentWindow covers enough days to compute today/week/month stats
	// without asking GitHub for a full year of contributionDays, which is
	// costly enough to occasionally trip GitHub's RESOURCE_LIMITS_EXCEEDED.
	recentWindow = 35 * 24 * time.Hour
)

type apiContributionRepository struct {
	configurationRepository configuration.ConfigurationRepository
	url                     string
}

func (repo *apiContributionRepository) GetContribution() (*contributions.Contribution, error) {
	token, err := repo.configurationRepository.GetToken()
	if err != nil {
		return nil, fmt.Errorf("getting token: %w", err)
	}

	username, err := repo.configurationRepository.GetUsername()
	if err != nil {
		return nil, fmt.Errorf("getting username: %w", err)
	}

	request, err := http.NewRequest("POST", repo.url, bytes.NewBuffer(payload(username)))
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}
	request.Header.Set("Authorization", "bearer "+*token)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("calling GitHub API: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("reading GitHub API response: %w", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API returned status %d: %s", response.StatusCode, string(body))
	}

	var data JSONData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("parsing GitHub API response: %w", err)
	}

	if len(data.Errors) > 0 {
		return nil, fmt.Errorf("GitHub API returned errors: %s", data.Errors[0].Message)
	}

	name := data.GetName()
	if name == "" {
		return nil, fmt.Errorf("unexpected GitHub API response, missing user name: %s", string(body))
	}

	today := data.GetDayContributions()
	week := data.GetWeekContributions()
	month := data.GetMonthContributions()
	year := data.GetYearContributions()

	return contributions.NewContribution(name, today, week, month, year), nil
}

func NewApiRepository(configurationRepository configuration.ConfigurationRepository) contributions.ContributionRepository {
	return &apiContributionRepository{
		configurationRepository: configurationRepository,
		url:                     url,
	}
}

func payload(username *string) []byte {
	now := time.Now().UTC()
	recentFrom := now.Add(-recentWindow).Format(time.RFC3339)
	yearFrom := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
	to := now.Format(time.RFC3339)

	return []byte(`
{
  "query": "query { user(login: \"` + *username + `\") { name yearStats: contributionsCollection(from: \"` + yearFrom + `\", to: \"` + to + `\") { contributionCalendar { totalContributions } } recentStats: contributionsCollection(from: \"` + recentFrom + `\", to: \"` + to + `\") { contributionCalendar { weeks { contributionDays { contributionCount date } } } } } }"
}
`)
}
