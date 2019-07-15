package world

import (
	"fmt"
)

// Monster ...
type Monster struct {
	ID      int
	City    string
	OldCity string
	Alive   bool
	Moves   int32
}

// MakeMonsters ...
func MakeMonsters(n int) map[int]Monster {
	monsters := make(map[int]Monster)

	for i := 1; i <= n; i = i + 1 {
		monsters[i] = Monster{
			ID:    i,
			Alive: true,
		}
	}

	return monsters
}

// Move ...
func (m *Monster) Move(city string) bool {
	m.Moves = m.Moves + 1
	m.OldCity = m.City
	m.City = city

	if m.Moves > 10000 {
		return true
	}

	return false
}

// Kill ...
func (m *Monster) Kill() {
	m.Alive = false
	fmt.Printf("Kill ... monster '%d' have been killed! - %d moves \n", m.ID, m.Moves)
}
