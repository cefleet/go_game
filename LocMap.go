package main
import "strconv"
var LocMap  = make(map[string]*Location)
func makeMap(){	
	col := 0
	for row := 0; row < 10; row++ {
		for ; col < 10; col++{
			LocMap[strconv.Itoa(row)+"-"+strconv.Itoa(col)] = &Location{
				row,
				col,
				nil,
				0,
				strconv.Itoa(row)+"-"+strconv.Itoa(col),
			}
		}
		col = 0;
	}
}