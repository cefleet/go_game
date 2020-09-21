package main
import (
	"strconv"
)

type Player struct {
	Mob
	inventory []*Item
}

func (player *Player) Move(dir string) *Result{
	//dir := "1"
	var key, status, message, dirname string
	if dir == "1"{
		dirname = "up"
		key = strconv.Itoa(player.on.x)+"-"+strconv.Itoa(player.on.y-1)
	} else if dir == "2"{
		dirname = "down"
		key = strconv.Itoa(player.on.x)+"-"+strconv.Itoa(player.on.y+1)
	} else if dir == "3"{
		dirname="left"
		key = strconv.Itoa(player.on.x-1)+"-"+strconv.Itoa(player.on.y)
	} else if dir == "4"{
		dirname = "right"
		key = strconv.Itoa(player.on.x+1)+"-"+strconv.Itoa(player.on.y)
	}

	if val, ok := LocMap[key]; ok {
		player.MoveTo(val)
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
	
	var player = new(Player)

	player.id = 1
	player.hp = 10
	player.MoveTo(LocMap["5-5"])

	return player
}