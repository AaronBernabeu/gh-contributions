package entrypoint

import (
	"encoding/json"
	"fmt"

	"github.com/Aaronidas/gh-contributions/internal/contributions"
	"github.com/spf13/cobra"
)

type CobraFn func(cmd *cobra.Command, args []string)

func InitContributionsCmd(repository contributions.ContributionRepository) *cobra.Command {
	contributionsCmd := &cobra.Command{
		Use:   "contributions",
		Short: "Print your contributions for today and for the last 365 days",
		Run:   runContributionsCmd(repository),
	}

	return contributionsCmd
}

func runContributionsCmd(repository contributions.ContributionRepository) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		contribData, _ := repository.GetContribution()
		jsonData, _ := json.MarshalIndent(contribData, "", "  ")
		fmt.Println(string(jsonData))
	}
}
