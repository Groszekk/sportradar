package score_board

import "sportradar/models"

type ScoreBoard struct {
	matches []*models.Match
}

func (sb *ScoreBoard) StartGame(homeTeam, awayTeam string) {
	match := &models.Match{
		HomeTeam:  homeTeam,
		AwayTeam:  awayTeam,
		HomeScore: 0,
		AwayScore: 0,
	}
	sb.matches = append(sb.matches, match)
}

func (sb *ScoreBoard) FinishGame(homeTeam, awayTeam string) {
	var index int
	for i, match := range sb.matches {
		if match.HomeTeam == homeTeam && match.AwayTeam == awayTeam {
			index = i
			break
		}
	}
	sb.matches = append(sb.matches[:index], sb.matches[index+1:]...)
}
