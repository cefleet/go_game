package main

type Entity struct {
	on *Location
	id int
}

func (entity *Entity) MoveTo(location *Location) {
	if entity.on != nil {
		entity.on.Empty()
	}
	entity.on = location
	location.on = entity
	location.status = 1
}

func (entity *Entity) RemoveEntityFromBoard(){
	if entity.on != nil {
		entity.on.on = nil
	}
	entity.on = nil
}