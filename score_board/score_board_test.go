package score_board

import (
	"sportradar/models"
	"testing"
)

func TestStartGame(t *testing.T) {
	sb := ScoreBoard{}
	sb.StartGame("TeamA", "TeamB")

	summary := sb.Summary()
	if len(summary) != 1 {
		t.Errorf("Summary should contain one match after starting a game, got: %d", len(summary))
	}
}

func TestFinishGame(t *testing.T) {
	sb := ScoreBoard{}
	sb.StartGame("TeamA", "TeamB")
	sb.FinishGame("TeamA", "TeamB")

	summary := sb.Summary()
	if len(summary) != 0 {
		t.Errorf("Summary should be empty after finishing the game, got: %d", len(summary))
	}
}

func TestUpdateScore(t *testing.T) {
	sb := ScoreBoard{}
	sb.StartGame("TeamA", "TeamB")
	sb.UpdateScore("TeamA", "TeamB", 2, 1)

	summary := sb.Summary()
	if summary[0].HomeScore != 2 {
		t.Errorf("Home team score should be updated, got: %d", summary[0].HomeScore)
	}
	if summary[0].AwayScore != 1 {
		t.Errorf("Away team score should be updated, got: %d", summary[0].AwayScore)
	}
}

func TestSummary(t *testing.T) {
	scoreBoard := ScoreBoard{}

	scoreBoard.StartGame("Mexico", "Canada")
	scoreBoard.UpdateScore("Mexico", "Canada", 0, 5)

	scoreBoard.StartGame("Spain", "Brazil")
	scoreBoard.UpdateScore("Spain", "Brazil", 10, 2)

	scoreBoard.StartGame("Germany", "France")
	scoreBoard.UpdateScore("Germany", "France", 2, 2)

	scoreBoard.StartGame("Uruguay", "Italy")
	scoreBoard.UpdateScore("Uruguay", "Italy", 6, 6)

	scoreBoard.StartGame("Argentina", "Australia")
	scoreBoard.UpdateScore("Argentina", "Australia", 3, 1)

	expectedSummary := []*models.Match{
		{HomeTeam: "Uruguay", AwayTeam: "Italy", HomeScore: 6, AwayScore: 6},
		{HomeTeam: "Spain", AwayTeam: "Brazil", HomeScore: 10, AwayScore: 2},
		{HomeTeam: "Mexico", AwayTeam: "Canada", HomeScore: 0, AwayScore: 5},
		{HomeTeam: "Argentina", AwayTeam: "Australia", HomeScore: 3, AwayScore: 1},
		{HomeTeam: "Germany", AwayTeam: "France", HomeScore: 2, AwayScore: 2},
	}

	summary := scoreBoard.Summary()

	for i := range expectedSummary {
		if summary[i].HomeTeam != expectedSummary[i].HomeTeam ||
			summary[i].AwayTeam != expectedSummary[i].AwayTeam ||
			summary[i].HomeScore != expectedSummary[i].HomeScore ||
			summary[i].AwayScore != expectedSummary[i].AwayScore {
			t.Errorf("Expected summary does not match at index %d", i)
		}
	}
}
