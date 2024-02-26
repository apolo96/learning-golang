package p1

import (
	"io"
	"os"
	"sort"
)

type Ranker interface {
	Ranking() []string
}

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(teamX, teamY string, scoreX, scoreY int) {
	if scoreX == scoreY{
		return
	}
	if _, ok := l.Teams[teamX]; !ok {
		return
	}
	if _, ok := l.Teams[teamY]; !ok {
		return
	}
	var winner = teamX
	if scoreY > scoreX {
		winner = teamY
	}
	l.Wins[winner] += 1
}

func (l *League) Ranking() []string {
	winners := make([]string,0, len(l.Teams))
	for key := range l.Wins {
		winners = append(winners, key)
	}
	sort.Slice(winners, func(i, j int) bool {
		return l.Wins[winners[i]] > l.Wins[winners[j]]
	})
	return winners
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, v := range r.Ranking() {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func New(){
	l := &League{
		Name: "Polo League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", "France", 50, 70) // France
	l.MatchResult("India", "Nigeria", 85, 80) // India
	l.MatchResult("Italy", "India", 60, 55) // Italy
	l.MatchResult("France", "Nigeria", 100, 110) // Nigeria
	l.MatchResult("Italy", "Nigeria", 65, 70) // Nigeria
	l.MatchResult("France", "India", 95, 80) // France
	RankPrinter(l, os.Stdout)
}
