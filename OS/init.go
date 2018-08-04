package main

import (
	"fmt"
	"os/exec"
)

//I am going to add input here, but for testing purposes I have not
func main(){
	var rpass = "root"
	var password ="root"
	if password != rpass{
		fmt.Println("You suck")
		}else if password == rpass{
			root()
	}
}

//will read qpuroot.py, via a sh work around
func root(){
	//working on it
}

//will read qpu standard.py
func standard(){
	//run qpustandard
}