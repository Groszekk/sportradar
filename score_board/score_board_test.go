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
