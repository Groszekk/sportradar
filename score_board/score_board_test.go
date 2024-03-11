package score_board

import (
	"fmt"
	"reflect"
	"sportradar/models"
	"testing"
)

func TestStartGame(t *testing.T) {
	sb := NewScoreBoard()
	homeTeam := "TeamA"
	awayTeam := "TeamB"
	key := fmt.Sprintf("%s-%s", homeTeam, awayTeam)

	sb.StartGame(homeTeam, awayTeam)

	match, ok := sb.matches[key]
	if !ok {
		t.Errorf("Match should have key: %s", key)
	}

	if match.HomeScore > 0 || match.AwayScore > 0 {
		t.Error("Match should start with zero scores")
	}

	if len(sb.matches) != 1 {
		t.Errorf("Summary should contain one match after starting a game, got: %d", len(sb.matches))
	}
}

func TestUpdateScore(t *testing.T) {
	sb := NewScoreBoard()
	homeTeam := "TeamA"
	awayTeam := "TeamB"
	homeScore := 1
	awayScore := 2
	key := fmt.Sprintf("%s-%s", homeTeam, awayTeam)

	sb.StartGame(homeTeam, awayTeam)
	sb.UpdateScore(homeTeam, awayTeam, homeScore, awayScore)

	if sb.matches[key].HomeScore != homeScore {
		t.Errorf("Home team score should be updated, got: %d", sb.matches[key].HomeScore)
	}

	if sb.matches[key].AwayScore != awayScore {
		t.Errorf("Away team score should be updated, got: %d", sb.matches[key].AwayScore)
	}
}

func TestFinishGame(t *testing.T) {
	sb := NewScoreBoard()
	homeTeam := "TeamA"
	awayTeam := "TeamB"

	sb.FinishGame(homeTeam, awayTeam)
	if len(sb.matches) != 0 {
		t.Errorf("Summary should be empty after finishing the game, got: %d", len(sb.matches))
	}
}

func TestSummary(t *testing.T) {
	sb := NewScoreBoard()

	homeTeam1 := "TeamA"
	awayTeam1 := "TeamB"

	homeTeam2 := "TeamC"
	awayTeam2 := "TeamD"

	homeTeam3 := "TeamE"
	awayTeam3 := "TeamF"

	sb.StartGame(homeTeam1, awayTeam1)
	sb.StartGame(homeTeam2, awayTeam2)
	sb.StartGame(homeTeam3, awayTeam3)

	sb.UpdateScore(homeTeam1, awayTeam1, 2, 2)
	sb.UpdateScore(homeTeam2, awayTeam2, 3, 0)
	sb.UpdateScore(homeTeam3, awayTeam3, 1, 5)

	expectedSummary := []*models.Match{
		{HomeTeam: homeTeam3, AwayTeam: awayTeam3, HomeScore: 1, AwayScore: 5},
		{HomeTeam: homeTeam1, AwayTeam: awayTeam1, HomeScore: 2, AwayScore: 2},
		{HomeTeam: homeTeam2, AwayTeam: awayTeam2, HomeScore: 3, AwayScore: 0},
	}

	summary := sb.Summary()

	if !reflect.DeepEqual(summary, expectedSummary) {
		t.Error("Summary did not return the expected result")
	}
}
