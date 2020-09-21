package main

type Entity struct {
	on *Location
	id int
}

func (entity *Entity) SetOn(location *Location) {
	entity.on = location
}

func (entity *Entity) RemoveEntityFromBoard(){
	if entity.on != nil {
		entity.on.on = nil
	}
	entity.on = nil
}