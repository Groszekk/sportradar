package main

import (
	"fmt"
	"sportradar/score_board"
)

func main() {
	// example output
	sb := score_board.NewScoreBoard()

	sb.StartGame("Mexico", "Canada")
	sb.UpdateScore("Mexico", "Canada", 0, 5)

	sb.StartGame("Spain", "Brazil")
	sb.UpdateScore("Spain", "Brazil", 10, 2)

	sb.StartGame("Germany", "France")
	sb.UpdateScore("Germany", "France", 2, 2)

	sb.StartGame("Uruguay", "Italy")
	sb.UpdateScore("Uruguay", "Italy", 6, 6)

	sb.StartGame("Argentina", "Australia")
	sb.UpdateScore("Argentina", "Australia", 3, 1)

	summary := sb.Summary()
	fmt.Println(summary)

	println("------------------------------------")
	for _, v := range summary {
		println("Home Team:", v.HomeTeam, "Score:", v.HomeScore)
		println("Away Team:", v.AwayTeam, "Score:", v.AwayScore)
		println("------------------------------------")
	}
}
