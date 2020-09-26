package main
import "strconv"

type gameMap struct {
	grid map[string]*Location
}

//TODO ADD Functions On The MAP HERE if needed
var Map = new(gameMap)

func setupMap() {
	grid := make(map[string]*Location)

	col := 0
	for row := 0; row < 10; row++ {
		for ; col < 10; col++{
			grid[strconv.Itoa(row)+"-"+strconv.Itoa(col)] = &Location{
				row,
				col,
				nil,
				0,
				strconv.Itoa(row)+"-"+strconv.Itoa(col),
			}
		}
		col = 0;
	}
	Map.grid = grid
}