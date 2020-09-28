package main
import "strconv"
type Mob struct {
	Entity
	hp int
	name string
	inventory []*Item
}

func (mob *Entity) Move(dir string) *Result{
	var key, status, message string
	//not loving this.
	if dir == "up"{
		key = strconv.Itoa(mob.on.x)+"-"+strconv.Itoa(mob.on.y-1)
	} else if dir == "down"{
		key = strconv.Itoa(mob.on.x)+"-"+strconv.Itoa(mob.on.y+1)
	} else if dir == "left"{
		key = strconv.Itoa(mob.on.x-1)+"-"+strconv.Itoa(mob.on.y)
	} else if dir == "right"{
		key = strconv.Itoa(mob.on.x+1)+"-"+strconv.Itoa(mob.on.y)
	}

	if val, ok := Map.grid[key]; ok {
		mob.MoveTo(val)
		status ="complete"
		message = "Moved "+dir+" one space."
	} else {
		status = "not-possible"
		message = "Space is outside of the map."
	}	

	return &Result{status,message,nil}
}