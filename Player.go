package main

var Player = new(Mob)

func setupPlayer(){
	Player.id = 1
	Player.hp = 10
	Player.MoveTo(Map.grid["5-5"])
}