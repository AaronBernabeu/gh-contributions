package contributions

type ContributionRepository interface {
	GetContribution() (*Contribution, error)
}
