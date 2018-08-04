package main

import (
	"fmt"
	"os/exec"
)

func main(){
	var rpass = "root"
	var password ="root"
	if password != rpass{
		fmt.Println("You suck")
		}else if password == rpass{
			root()
	}
}

func root(){
	//working on it
}

func standard(){
	//run qpustandard
}