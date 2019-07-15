package engine

import (
	"fmt"

	"github.com/car12o/monsters/world"
)

// Engine ...
type Engine struct {
	worldMap      world.Map
	aliveMonsters map[int]world.Monster
	deadMonsters  map[int]world.Monster
}

// New ...
func New(worldMap world.Map, monsters map[int]world.Monster) Engine {
	return Engine{
		worldMap:      worldMap,
		aliveMonsters: monsters,
		deadMonsters:  make(map[int]world.Monster),
	}
}

// StartGame ...
func (eng *Engine) StartGame() {
	for {
		for i, monster := range eng.aliveMonsters {
			worldMap := (*eng).worldMap
			aliveMonsters := (*eng).aliveMonsters

			city := eng.worldMap.GetMonsterNextCity(monster.City)
			if city.Name == "" {
				keys := eng.killMonsters(map[int]world.Monster{monster.ID: monster})
				err := fmt.Sprintf("by monster %d [no more links to go]", keys[0])
				worldMap.DestroyCity(monster.City, err)
				continue
			}

			destroyMonster := monster.Move(city.Name)
			if destroyMonster {
				eng.killMonsters(map[int]world.Monster{monster.ID: monster})
				worldMap.RemoveMonsterFromCity(aliveMonsters[i])
				continue
			}
			aliveMonsters[i] = monster

			destroyCity := eng.worldMap.MoveMonster(aliveMonsters[i], city.Name)
			if destroyCity {
				keys := eng.killMonsters(worldMap[city.Name].Monsters)
				err := fmt.Sprintf("by monsters '%d' and '%d'", keys[0], keys[1])
				worldMap.DestroyCity(city.Name, err)
			}
		}

		if len(eng.aliveMonsters) == 0 {
			break
		}
	}

	eng.worldMap.Print()
}

func (eng *Engine) killMonsters(monsters map[int]world.Monster) []int {
	aliveMonsters := (*eng).aliveMonsters
	deadMonsters := (*eng).deadMonsters

	var keys []int
	for key, monster := range monsters {
		keys = append(keys, key)
		delete(aliveMonsters, monster.ID)
		monster.Kill()
		deadMonsters[monster.ID] = monster
	}

	return keys
}
