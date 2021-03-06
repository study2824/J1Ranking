package action

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type TeamData struct {
	TeamName       string `json:"team_name"`
	Rank           string `json:"rank"`
	Points         string `json:"points"`
	MatchPlayed    string `json:"match_played"`
	Win            string `json:"win"`
	Draw           string `json:"draw"`
	Lose           string `json:"lose"`
	GoalScored     string `json:"goal_scored"`
	GoalConceded   string `json:"goal_conceded"`
	GoalDifference string `json:"goal_difference"`
	GoalAve        string `json:"goal_ave"`
	ConcededAve    string `json:"conceded_ave"`
}

func getData(league string, year string) ([]TeamData, error) {
	var teams []TeamData
	i, err := strconv.Atoi(year)
	if i <= 2011 {
		return nil, errors.New("データが対応していません")
	}

	url := "https://www.football-lab.jp/summary/team_ranking/" + league + "/?year=" + year
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	doc.Find("table#standing > tbody > tr").Each(func(i int, s *goquery.Selection) {
		var team TeamData
		s.Find("td").Each(func(i int, v *goquery.Selection) {
			switch i {
			case 0:
				team.Rank = v.Text()
			case 1:
				return //チームエンブレム
			case 2:
				team.TeamName = v.Find("span.dsktp").Text()
			case 3:
				team.Points = v.Text()
			case 4:
				team.MatchPlayed = v.Text()
			case 5:
				team.Win = v.Text()
			case 6:
				team.Draw = v.Text()
			case 7:
				team.Lose = v.Text()
			case 8:
				team.GoalScored = v.Text()
			case 9:
				team.GoalConceded = v.Text()
			case 10:
				team.GoalDifference = v.Text()
			case 11:
				team.GoalAve = v.Text()
			case 12:
				team.ConcededAve = v.Text()
			}
		})
		teams = append(teams, team)
	})
	return teams, err
}

func Ranking(league string, year string) ([]TeamData, error) {
	return getData(league, year)
}

func TeamDetail(league string, year string, n string) (TeamData, error) {
	teams, err := getData(league, year)
	if err != nil {
		return TeamData{}, err
	}

	for i := 0; i < len(teams); i++ {
		if teams[i].TeamName == n {
			return teams[i], err
		}
	}
	return TeamData{}, err
}

func ScoreDifference(league string, year string, team1Name string, team2Name string) (TeamData, TeamData, int, error) {
	teams, err := getData(league, year)
	if err != nil {
		return TeamData{}, TeamData{}, 0, err
	}

	var team1 TeamData
	var team2 TeamData

	for i := 0; i < len(teams); i++ {
		if teams[i].TeamName == team1Name {
			team1 = teams[i]
		}
		if teams[i].TeamName == team2Name {
			team2 = teams[i]
		}
	}

	team1Point, err := strconv.Atoi(team1.Points)
	team2Point, err := strconv.Atoi(team2.Points)

	return team1, team2, team1Point - team2Point, nil
}
