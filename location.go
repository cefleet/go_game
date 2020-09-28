package main

type Location struct{
	x int
	y int
	on *Entity
	status int
	name string
}

func (location *Location) SetOn(entity *Entity) {
	location.on = entity
}

func (location *Location) Empty(){
	location.on = nil
}