package main
import "strconv"
type Mob struct {
	Entity
	hp int
	name string
	inventory []*Item
}

func (mob *Entity) MoveTo(location *Location){
	mob.SetOn(location)
	location.SetOn(mob)
}

func (mob *Entity) Move(dir string) *Result{
	var key, status, message, dirname string
	//not loving this.
	if dir == "1"{
		dirname = "up"
		key = strconv.Itoa(mob.on.x)+"-"+strconv.Itoa(mob.on.y-1)
	} else if dir == "2"{
		dirname = "down"
		key = strconv.Itoa(mob.on.x)+"-"+strconv.Itoa(mob.on.y+1)
	} else if dir == "3"{
		dirname="left"
		key = strconv.Itoa(mob.on.x-1)+"-"+strconv.Itoa(mob.on.y)
	} else if dir == "4"{
		dirname = "right"
		key = strconv.Itoa(mob.on.x+1)+"-"+strconv.Itoa(mob.on.y)
	}

	if val, ok := Map.grid[key]; ok {
		mob.MoveTo(val)
		status ="complete"
		message = "Moved "+dirname+" one space."
	} else {
		status = "not-possible"
		message = "Space is outside of the map."
	}	

	return &Result{
		status,
		message,
	}
}