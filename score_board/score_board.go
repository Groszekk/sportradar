package score_board

import (
	"sort"
	"sportradar/models"
)

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

func (sb *ScoreBoard) Summary() []*models.Match {
	// Introsort
	// Quicksort, Heapsort and Insertionsort
	sort.SliceStable(sb.matches, func(i, j int) bool {
		totalScore1 := sb.matches[i].HomeScore + sb.matches[i].AwayScore
		totalScore2 := sb.matches[j].HomeScore + sb.matches[j].AwayScore
		if totalScore1 == totalScore2 {
			return i > j
		}
		return totalScore1 > totalScore2
	})
	return sb.matches
}
