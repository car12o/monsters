package main

import (
	"flag"
	"strconv"

	"github.com/car12o/monsters/engine"
	"github.com/car12o/monsters/world"
)

func main() {
	monsters := flag.String("m", "10", "number of monsters")
	file := flag.String("f", "./files/world_map_small.txt", "path for world map file")
	flag.Parse()

	n, err := strconv.Atoi(*monsters)
	if err != nil {
		panic(err)
	}

	worldMap, err := world.LoadMap(*file)
	if err != nil {
		panic(err)
	}

	monster := world.MakeMonsters(n)

	eng := engine.New(worldMap, monster)

	eng.StartGame()
}
