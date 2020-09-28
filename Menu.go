package main

import (
	"fmt"
	//"strconv"
)

type MenuSelection struct{
	text string
	acceptedValue string
	valueOnAccept string //this is the value put in the arg
	action string
	actionOn Node
}

type MenuScreen struct {
	name string
	text string
	input bool
	selections []MenuSelection
}

func (menu *MenuScreen) AccecptValue(value string)*Result{
	for _, item := range menu.selections { 
		if item.acceptedValue == value {
			res := doAction(item.action, item.valueOnAccept, item.actionOn)
			return res
		}
	}
	return &Result{"fail", "No-Menu-Item",nil}
}

func (menu *MenuScreen) RunMenu()*Result{
	var res string
	PrintOut(menu.text+"\n")
	for _, opt := range menu.selections{
		PrintOut(" ^g^^bld^^rvs^ "+opt.acceptedValue+". "+opt.text+" ^e^ ")
	}
	if menu.input {
		fmt.Scanln(&res)
		return menu.AccecptValue(res)
	}
	return &Result{"success", "menu-printed", nil}
}

type GameMenu struct{
	menus map[string]*MenuScreen
} 

func (menu *GameMenu) RunMenu(menuScreen string) *Result{
	var res *Result
	if val, ok := menu.menus[menuScreen]; ok {
		res = val.RunMenu()
	} else {
		res = &Result{"error", "key-not-found", nil}
	}
	return res
}

var Menu = new(GameMenu)

func setupMenus(){

	MenuScreens := make(map[string]*MenuScreen)
	start := new(MenuScreen)
	main := new(MenuScreen)
	move := new(MenuScreen)

	start.name = "start"
	start.text = `^y^^bB^^bld^Welcome^e^^y^^bB^ to the ^udl^ game .`
	start.input = false
	
	main.name = "main"
	main.text = `What would you like to do?`

	main.input = true
	main.selections = []MenuSelection{
		{"Move","1","move", "RunMenu",Menu},
	}	
	
	move.name = "move"
	move.text =`What direction would you like to move?`

	move.input = true
	move.selections = []MenuSelection{
		{"Up","1","up","Move",Player},
		{"Down","2","down","Move",Player},
		{"Left","3","left","Move",Player},
		{"Right","4","right","Move",Player},
	}

	MenuScreens["start"] = start
	MenuScreens["main"] = main
	MenuScreens["move"] = move

	Menu.menus = MenuScreens
}