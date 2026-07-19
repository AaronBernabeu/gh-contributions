package repository

import "time"

type ContributionDay struct {
	ContributionCount int    `json:"contributionCount"`
	Date              string `json:"date"`
}

type ContributionWeek struct {
	ContributionDays []ContributionDay `json:"contributionDays"`
}

type ContributionCalendar struct {
	TotalContributions int                `json:"totalContributions"`
	Weeks              []ContributionWeek `json:"weeks"`
}

type ContributionsCollection struct {
	ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
}

type User struct {
	Name        string                  `json:"name"`
	YearStats   ContributionsCollection `json:"yearStats"`
	RecentStats ContributionsCollection `json:"recentStats"`
}

type Data struct {
	User User `json:"user"`
}

type GraphQLError struct {
	Message string `json:"message"`
}

type JSONData struct {
	Data   Data           `json:"data"`
	Errors []GraphQLError `json:"errors"`
}

func (d JSONData) GetName() string {
	return d.Data.User.Name
}

func (d JSONData) GetWeekContributions() int {
	today := time.Now().Format("2006-01-02")
	for _, week := range d.Data.User.RecentStats.ContributionCalendar.Weeks {
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
	for _, week := range d.Data.User.RecentStats.ContributionCalendar.Weeks {
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

	for _, week := range d.Data.User.RecentStats.ContributionCalendar.Weeks {
		for _, day := range week.ContributionDays {
			curTime, _ := time.Parse("2006-01-02", day.Date)
			if curTime.Month() == now.Month() && curTime.Year() == now.Year() {
				count = count + day.ContributionCount
			}
		}
	}

	return
}

func (d JSONData) GetYearContributions() int {
	return d.Data.User.YearStats.ContributionCalendar.TotalContributions
}

func (w *ContributionWeek) getContributionCount() (count int) {
	count = 0
	for _, day := range w.ContributionDays {
		count = count + day.ContributionCount
	}

	return
}
