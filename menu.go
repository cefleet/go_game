package main

import (
	"fmt"
)

type MenuSelection struct{
	name string
	acceptedValue string
	action string
	actionOn Node
}

type Menu struct {
	name string
	text string
	input bool
	selections []MenuSelection
}

func (menu *Menu) AccecptValue(value string)*Result{
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

func (menu *Menu) RunMenu()*Result{
	var res string
	PrintOut(menu.text)
	if menu.input {
		fmt.Scanln(&res)
		menu.AccecptValue(res)
	}
	return &Result{"sucess", "menu-printed"}
}

var Menus = make(map[string]*Menu)

func makeMenus(){

	start := new(Menu)
	main := new(Menu)
	move := new(Menu)

	start.name = "start"
	start.text = `|y||bB||bld|Welcome|e||y||bB| to the |udl| game .`
	start.input = false
	
	main.name = "main"
	main.text = `
What would you like to do?
|g||rvs|1.|e| Move
|g||rvs|2.|e| Look
|g||rvs|3.|e| Inventory
|g||rvs|4.|e| End
`
	main.input = true
	main.selections = []MenuSelection{
		{"move","1","RunMenu",move},
	}	
	
	move.name = "move"
	move.text =`
|g||rvs|1.|e| Up
|g||rvs|2.|e| Down
|g||rvs|3.|e| Left
|g||rvs|4.|e| Right
|r||rvs|0.|e| Go Back
`

move.input = true
move.selections = []MenuSelection{
	{"Up","1","Move",player},
	{"Down","2","Move",player},
	{"Left","3","Move",player},
	{"Right","4","Move",player},
}

	Menus["start"] = start
	Menus["main"] = main
	Menus["move"] = move

}
