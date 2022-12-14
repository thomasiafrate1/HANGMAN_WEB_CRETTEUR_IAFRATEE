package piscine

import "sort"

type Score struct {
	Name   string
	Points int
}

func ScoreJoueur(family []Score) []Score {
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Points > family[j].Points
	})
	if len(family) > 3 {
		family = family[:3]
	}
	return family
}
