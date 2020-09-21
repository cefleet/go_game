package main

type Mob struct {
	Entity
	hp int
}

func (mob *Entity) MoveTo(location *Location){
	mob.SetOn(location)
	location.SetOn(mob)
}