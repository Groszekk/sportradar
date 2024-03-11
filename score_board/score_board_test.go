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
