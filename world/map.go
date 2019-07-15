package world

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// City ...
type City struct {
	Name     string
	Links    map[string]string
	Monsters map[int]Monster
}

// Map ...
type Map map[string]City

// LoadMap ...
func LoadMap(filePath string) (Map, error) {
	worldMap := make(map[string]City)

	file, err := os.Open(filePath)
	if err != nil {
		return Map{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")

		city := City{
			Name:     data[0],
			Links:    make(map[string]string),
			Monsters: make(map[int]Monster),
		}

		for _, val := range data[1:] {
			split := strings.Split(val, "=")
			city.Links[split[0]] = split[1]
		}

		if _, ok := worldMap[city.Name]; ok {
			return Map{}, fmt.Errorf("Invalid source map file, repeated city '%s'", city.Name)
		}
		worldMap[city.Name] = city
	}

	if err := scanner.Err(); err != nil {
		return Map{}, err
	}

	return worldMap, nil
}

// GetMonsterNextCity ...
func (m *Map) GetMonsterNextCity(cityName string) City {
	var result City

	if cityName == "" {
		return m.getRandomCity()
	}

	city := (*m)[cityName]
	if len(city.Links) == 0 {
		return result
	}

	rand := 0
	if len(city.Links) > 1 {
		rand = randomize(len(city.Links) - 1)
	}

	i := 0
	for key := range city.Links {
		if i == rand {
			result = (*m)[city.Links[key]]
		}
		i++
	}

	return result
}

// MoveMonster ...
func (m *Map) MoveMonster(monster Monster, cityName string) bool {
	m.RemoveMonsterFromCity(monster)

	city := (*m)[cityName]

	if city.Monsters == nil {
		city.Monsters = make(map[int]Monster)
	}
	city.Monsters[monster.ID] = monster
	if len(city.Monsters) > 1 {
		return true
	}

	return false
}

// RemoveMonsterFromCity ...
func (m *Map) RemoveMonsterFromCity(monster Monster) {
	if _, ok := (*m)[monster.OldCity]; ok {
		delete((*m)[monster.OldCity].Monsters, monster.ID)
	}
}

// DestroyCity ...
func (m *Map) DestroyCity(cityName string, err string) {
	worldMap := (*m)
	delete(worldMap, cityName)

	for mapKey, city := range worldMap {
		for cityKey, link := range city.Links {
			if cityName == link {
				delete(worldMap[mapKey].Links, cityKey)
			}
		}
	}

	fmt.Printf("BOOM ... city '%s' was destroyed! %s \n", cityName, err)
}

// Print ...
func (m *Map) Print() {
	var file []string

	for _, city := range *m {
		line := fmt.Sprintf("%s", city.Name)
		for key, val := range city.Links {
			line = fmt.Sprintf("%s %s=%s", line, key, val)
		}
		file = append(file, line)
	}

	fmt.Println("### Remaining World Map ###")
	for _, line := range file {
		fmt.Println(line)
	}
}

func (m *Map) getRandomCity() City {
	var result City

	rand := randomize(len(*m) - 1)

	i := 0
	for key := range *m {
		if i == rand {
			result = (*m)[key]
		}
		i++
	}

	return result
}

func randomize(max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(max)
}
