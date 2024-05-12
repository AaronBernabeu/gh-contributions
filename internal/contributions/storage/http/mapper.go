package repository

import "time"

type ContributionDay struct {
	Color             string `json:"color"`
	ContributionCount int    `json:"contributionCount"`
	Date              string `json:"date"`
	Weekday           int    `json:"weekday"`
}

type ContributionWeek struct {
	ContributionDays []ContributionDay `json:"contributionDays"`
	FirstDay         string            `json:"firstDay"`
}

type ContributionCalendar struct {
	Colors             []string           `json:"colors"`
	TotalContributions int                `json:"totalContributions"`
	Weeks              []ContributionWeek `json:"weeks"`
}

type ContributionsCollection struct {
	ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
}

type User struct {
	Name                    string                  `json:"name"`
	ContributionsCollection ContributionsCollection `json:"contributionsCollection"`
}

type Data struct {
	User User `json:"user"`
}

type JSONData struct {
	Data Data `json:"data"`
}

func (d JSONData) GetName() string {
	return d.Data.User.Name
}

func (d JSONData) GetWeekContributions() int {
	today := time.Now().Format("2006-01-02")
	for _, week := range d.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			if day.Date == today {
				return week.getContributionCount()
			}
		}
	}

	return 0
}

func (d JSONData) GetDayContributions() int {
	today := time.Now().Format("2006-01-02")
	for _, week := range d.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			if day.Date == today {
				return day.ContributionCount
			}
		}
	}

	return 0
}

func (d JSONData) GetMonthContributions() (count int) {
	now := time.Now()
	count = 0

	for _, week := range d.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			curTime, _ := time.Parse("2006-01-02", day.Date)
			if curTime.Month() == now.Month() && curTime.Year() == now.Year() {
				count = count + day.ContributionCount
			}
		}
	}

	return
}

func (d JSONData) GetTotalContributions() int {
	return d.Data.User.ContributionsCollection.ContributionCalendar.TotalContributions
}

func (w *ContributionWeek) getContributionCount() (count int) {
	count = 0
	for _, day := range w.ContributionDays {
		count = count + day.ContributionCount
	}

	return
}
