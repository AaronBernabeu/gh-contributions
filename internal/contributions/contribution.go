package contributions

type Contribution struct {
	Name  string `json:"name"`
	Today int    `json:"today"`
	Week  int    `json:"week"`
	Month int    `json:"month"`
	Year  int    `json:"year"`
}

func NewContribution(Name string, Today int, Week int, Month int, Year int) (contribution *Contribution) {
	contribution = &Contribution{
		Name:  Name,
		Today: Today,
		Week:  Week,
		Month: Month,
		Year:  Year,
	}

	return
}
