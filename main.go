package main
import (
	"reflect"
	//"fmt"
	"strconv"
)

type Result struct{
	status string
	message string
	linesOut []string
}

var playing bool
//var player *Player
//var PlayerActions = make(map[string]func(string)*Result)

func main(){
	setupMap()
	setupPlayer()
	setupMenus()
	gameLoop()
	//asciichart()		
}

//The actions need to be "done ON" specific items that are ready to have this action done on it.
func doAction(action string, value string, actionOn Node) *Result{
	//This should be 1 or 0. The aregument can just simply call something without an arg sometimes.
	len := 1
	if value == "" {
		len = 0
	}
	inputs := make([]reflect.Value, len) //The call requirers an array of reflect.Values?!?
	
	//do not need it if the function is to be called without an argument
	if len > 0 {
		inputs[0] = reflect.ValueOf(value)//but they add the args to the function
	}

	//Everything that is supposed to be called as an action needs to return a result
	result := reflect.ValueOf(actionOn).MethodByName(action).Call(inputs)[0].Interface().(*Result)
	return result

}




func gameLoop(){
	playing = true
	Menu.menus["start"].RunMenu()

	//same as an while loop in js and python
	for playing {
		CallClear()
		mapRes := Map.DrawMap()
		menuItems := []string{
			"",
			"	| HP 		: "+strconv.Itoa(Player.hp),
			"	| Tile 		: "+Player.on.name,
			"	| Inv/count	: "+strconv.Itoa(len(Player.inventory)),
		}

		for i := 0; i < 12; i++{
			mapOut := ""
			menuOut := ""
			if len(mapRes.linesOut) > i {
				mapOut = mapRes.linesOut[i]
			}
			if len(menuItems) > i {
				menuOut = menuItems[i]
			}

			PrintOut(mapOut+menuOut)
			PrintOut("\n")
		}

		Menu.menus["main"].RunMenu()
	}
}