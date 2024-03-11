package score_board

import "testing"

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
