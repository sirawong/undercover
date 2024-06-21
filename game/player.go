package game

import (
	"math/rand"
	"time"
)

const (
	Undercover = "Undercover"
	Civilian   = "Civilian"
)

type Player struct {
	Name string
	Role string
}

func AssignRoles(players []Player, numUndercover, numCivilians int) []Player {
	roles := make([]string, len(players))
	for i := 0; i < numUndercover; i++ {
		roles[i] = Undercover
	}
	for i := numUndercover; i < numUndercover+numCivilians; i++ {
		roles[i] = Civilian
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(roles), func(i, j int) {
		roles[i], roles[j] = roles[j], roles[i]
	})
	for i := range players {
		players[i].Role = roles[i]
	}
	return players
}
