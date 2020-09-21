package main
import (
	"fmt"
	"strconv"
)

func StartMenu(){
	PrintOut("|y||bB||bld|Welcome|e||y||bB| to the |udl| game .")
}

func BeforeAction(player *Player){
	PrintOut("You are on location |r|"+player.on.name+".")
	PrintOut("You have |g||rvs|"+strconv.Itoa(player.hp)+"|e| HP.")
}

func PrintResult(result *Result){
	fmt.Println(result.status)
	fmt.Println(result.message)
}

func MoveMenu() string{
	var dir,confirm string
	PrintOut("What direction would you like to Move?")
	PrintOut(`
|g||rvs|1.|e| Up
|g||rvs|2.|e| Down
|g||rvs|3.|e| Left
|g||rvs|4.|e| Right
|r||rvs|0.|e| Go Back
`)
	fmt.Scanln(&dir)

	if dir == "1"{
		fmt.Println("You chose Up.")
	} else if dir == "2"{
		fmt.Println("You chose Up.")
	} else if dir == "3"{
		fmt.Println("You chose Left.")
	} else if dir == "4"{
		fmt.Println("You chose Right.")
	} else if dir == "0"{
		return "0"
	}
	fmt.Println("Are you sure?\n  1. Yes\n  2. No\n")
	fmt.Scanln(&confirm)

	if(confirm == "1"){
		return dir
	} else {
		return MoveMenu()
	}
}

func MainMenu() (action, value string ){
	var choice string
	fmt.Println("What whould you like to do?")
	fmt.Println("  1. Move\n  2. Look\n  3. Inventory\n  0. Exit")
	fmt.Scanln(&choice)

	if choice == "1"{
		action = "Move"
		value = MoveMenu()
		return action, value
	}

	return "none", "none"
}