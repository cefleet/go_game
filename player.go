package main
import (
	"strconv"
//	"fmt"
)

type Player struct {
	Mob
	inventory []*Item
}

func (p *Player) Move(dir string) *Result{
	//dir := "1"
	var key, status, message, dirname string
	if dir == "1"{
		dirname = "up"
		key = strconv.Itoa(p.on.x)+"-"+strconv.Itoa(p.on.y-1)
	} else if dir == "2"{
		dirname = "down"
		key = strconv.Itoa(p.on.x)+"-"+strconv.Itoa(p.on.y+1)
	} else if dir == "3"{
		dirname="left"
		key = strconv.Itoa(p.on.x-1)+"-"+strconv.Itoa(p.on.y)
	} else if dir == "4"{
		dirname = "right"
		key = strconv.Itoa(p.on.x+1)+"-"+strconv.Itoa(p.on.y)
	}

	if val, ok := LocMap[key]; ok {
		p.MoveTo(val)
		status ="complete"
		message = "You have moved "+dirname+" one space."
	} else {
		status = "not-possible"
		message = "That space is outside of the map."
	}
	

	return &Result{
		status,
		message,
	}
}

func setupPlayer() *Player{
	
	var p = new(Player)

	p.id = 1
	p.hp = 10
	p.MoveTo(LocMap["5-5"])

	return p
}