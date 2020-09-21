package main
import "reflect"

type Result struct{
	status string
	message string
}

var playing bool
var player *Player
var PlayerActions = make(map[string]func(string)*Result)

func main(){
	makeMap()
	player = setupPlayer()
	gameLoop()		
}

func doAction(action string, value string) *Result{
	inputs := make([]reflect.Value, 1) //no clue how these lines work
	inputs[0] = reflect.ValueOf(value)//but they add the args to the function

	resultArr := reflect.ValueOf(player).MethodByName(action).Call(inputs)
	return resultArr[0].Interface().(*Result)

	//return result
}

func gameLoop(){
	playing = true
	StartMenu()
	for playing {
		BeforeAction(player)
		action, value := MainMenu()
		result := doAction(action ,value)		
		
		PrintResult(result)
		
	}
}