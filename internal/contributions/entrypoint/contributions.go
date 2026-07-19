package entrypoint

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AaronBernabeu/gh-contributions/internal/contributions"
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
		contribData, err := repository.GetContribution()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		jsonData, err := json.MarshalIndent(contribData, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		fmt.Println(string(jsonData))
	}
}
