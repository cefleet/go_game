package main
import (
	"strconv"
	//"fmt"
	//"sort"
)

type gameMap struct {
	grid map[string]*Location
}

func (m *gameMap) DrawMap()*Result{
	param := 4
	lineArr := []string{}
	pX := Player.on.x
	pY := Player.on.y
	row := pY-param
	lineArr = append(lineArr, "")
	for ; row <= pY+param; row++{
		// topLine := ""
		// for col := pX-3; col <= pX+3; col++{
		// 	topLine += ReturnOut(" ___")
		// }
		// lineArr = append(lineArr, topLine)
		contentLine := ""
		for col := pX-param; col <= pX+param; col++{
			
			//if it is an actual item
			if res, ok := m.grid[strconv.Itoa(col)+"-"+strconv.Itoa(row)]; ok {
				if res.on != nil {
					contentLine += ReturnOut("^bB^ ^pl^ ")					
				} else if res.status ==1{
					contentLine += ReturnOut("^rB^   ")
				} else {
					contentLine += ReturnOut("^gB^   ")
				}
			} else {
				contentLine += ReturnOut("^wB^   ")
			}			
	   	}
		contentLine += ""
		lineArr = append(lineArr, contentLine)
	}

	// bottomLine := ""
	// for col := pX-3; col <= pX+3; col++{
	// 	bottomLine += ReturnOut(" ")
	// }
	//lineArr = append(lineArr, bottomLine)
	//lineArr = append(lineArr, "")

	//Proof Of Concept for printing from an array
	// for _, line := range lineArr{
	// 	PrintOut(line)
	// 	PrintOut("\n")

	// } 

	return &Result{"sucess","map-printed", lineArr}

}


var Map = new(gameMap)

func setupMap() {
	grid := make(map[string]*Location)

	row := 0
	for col := 0; col < 10; col++ {
		for ; row < 10; row++{
			grid[strconv.Itoa(col)+"-"+strconv.Itoa(row)] = &Location{
				col,
				row,
				nil,
				0,
				strconv.Itoa(col)+"-"+strconv.Itoa(row),
			}
		}
		row = 0;
	}
	Map.grid = grid
}