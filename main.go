package main
import (
	"reflect"
	"fmt"
)

type Result struct{
	status string
	message string
}

var playing bool
var player *Player
//var PlayerActions = make(map[string]func(string)*Result)

func main(){
	makeMap()
	player = setupPlayer()
	makeMenus()
	gameLoop()		
}

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
	fmt.Println(result)
	return result

}

func gameLoop(){
	playing = true
	Menus["start"].RunMenu()
	for playing {
		Menus["main"].RunMenu()		
	}
}