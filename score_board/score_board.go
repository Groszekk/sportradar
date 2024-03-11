package score_board

import (
	"sort"
	"sportradar/models"
)

type ScoreBoard struct {
	matches map[string]*models.Match
}

func NewScoreBoard() *ScoreBoard {
	return &ScoreBoard{
		matches: make(map[string]*models.Match),
	}
}

func (sb *ScoreBoard) StartGame(homeTeam, awayTeam string) {
	match := &models.Match{
		HomeTeam:  homeTeam,
		AwayTeam:  awayTeam,
		HomeScore: 0,
		AwayScore: 0,
	}
	key := homeTeam + "-" + awayTeam // key: homeTeam-awayTeam
	sb.matches[key] = match
}

func (sb *ScoreBoard) UpdateScore(homeTeam, awayTeam string, homeScore, awayScore int) {
	key := homeTeam + "-" + awayTeam
	match, ok := sb.matches[key]
	if ok {
		match.HomeScore = homeScore
		match.AwayScore = awayScore
	}
}

func (sb *ScoreBoard) FinishGame(homeTeam, awayTeam string) {
	key := homeTeam + "-" + awayTeam
	delete(sb.matches, key)
}

func (sb *ScoreBoard) Summary() []*models.Match {
	// Tworzymy pustą listę par klucz-wartość
	var sortedMatches []*models.Match

	// Przekształcamy mapę w listę
	for _, match := range sb.matches {
		sortedMatches = append(sortedMatches, match)
	}

	// Sortujemy listę meczów według sumy punktów
	sort.SliceStable(sortedMatches, func(i, j int) bool {
		totalScore1 := sortedMatches[i].HomeScore + sortedMatches[i].AwayScore
		totalScore2 := sortedMatches[j].HomeScore + sortedMatches[j].AwayScore

		// sortujemy malejąco
		return totalScore1 > totalScore2
	})

	return sortedMatches
}
