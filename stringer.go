package main

import (
	"strings"
	"fmt"
)

var code = make(map[string]string)

func init(){
	//color
	code["r"] = "\033[31m"
	code["g"] = "\033[32m"
	code["b"] = "\033[34m"
	code["y"] = "\033[33m"
	code["p"] = "\033[35m"
	code["w"] = "\033[37m"
	code["gr"] = "\033[38m"
	//bg
	code["rB"] = "\u001b[41m"
	code["gB"] = "\u001b[42m"
	code["bB"] = "\u001b[44m"
	code["yB"] = "\u001b[43m"
	code["pB"] = "\u001b[45m"
	code["wB"] = "\u001b[47m"
	code["grB"] = "\u001b[48m"

	//modifiers
	code["bld"] = "\u001b[1m"
	code["udl"] = "\u001b[4m"
	code["rvs"] = "\u001b[7m"
	code["udl"] = "\u0033[4m"
	//end
	code["e"] = "\033[0m"

	//ascii
	code["pl"] = "@"

}

func ReturnOut(value string)string{
	out := strings.Split(value, "^")
	for i, v := range out {
		if val, ok := code[v]; ok {
			out[i] = val
		}
	}
	strOut := strings.Join(out, "")

	strOut += "\u001b[0m"
	return strOut
}

func PrintOut(value string){		
	fmt.Print(ReturnOut(value))
}