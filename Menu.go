package main

import (
	"fmt"
)

type MenuSelection struct{
	text string
	acceptedValue string
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
	for _, item := range menu.selections { //_would be idx
		//fmt.Println(idx)//possible have a menu list and use the idx
		if item.acceptedValue == value {
			argValue := value
			if item.action == "RunMenu"{
				argValue = ""
			}
			res := doAction(item.action, argValue, item.actionOn)
			return res
		}
	}
	return &Result{"fail", "No-Menu-Item"}
}

func (menu *MenuScreen) RunMenu()*Result{
	var res string
	PrintOut(menu.text)
	for _, opt := range menu.selections{
		PrintOut("|g||rvs|"+opt.acceptedValue+".|e| "+opt.text)
	}
	if menu.input {
		fmt.Scanln(&res)
		res := menu.AccecptValue(res)
		return res
	}
	return &Result{"success", "menu-printed"}
}

type GameMenu struct{
	menus map[string]*MenuScreen
} 

var Menu = new(GameMenu)

func setupMenus(){

	MenuScreens := make(map[string]*MenuScreen)
	start := new(MenuScreen)
	main := new(MenuScreen)
	move := new(MenuScreen)

	start.name = "start"
	start.text = `|y||bB||bld|Welcome|e||y||bB| to the |udl| game .`
	start.input = false
	
	main.name = "main"
	main.text = `What would you like to do?`

	main.input = true
	main.selections = []MenuSelection{
		{"Move","1","RunMenu",move},
	}	
	
	move.name = "move"
	move.text =`What direction would you like to move?`

	move.input = true
	move.selections = []MenuSelection{
		{"Up","1","Move",Player},
		{"Down","2","Move",Player},
		{"Left","3","Move",Player},
		{"Right","4","Move",Player},
	}

	MenuScreens["start"] = start
	MenuScreens["main"] = main
	MenuScreens["move"] = move

	Menu.menus = MenuScreens
}