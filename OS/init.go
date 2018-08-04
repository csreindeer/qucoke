package main

import (
	"fmt"
	"io/ioutil"
)

//I am going to add input here, but for testing purposes I have not
//will create a for loop as an "engine"
func main(){
	var rpass = "root"
	var password ="root"
	if password != rpass{
		fmt.Println("You suck")
		}else if password == rpass{
			root()
	}
	output()
}

//will read qpuroot.py, via a sh work around
func root(){
	exec.Command("/bin/sh", "/home/mymainman/bigboy/git qucoke/qucoke/OS/workaround.sh")
}

//will read qpu standard.py
func standard(){
	//run qpustandard
}
func output(){
	o, err := ioutil.ReadFile("output.txt")
	if err != nil{
		fmt.Println(err)
	}
	str := string(o)

	fmt.Println(str)
}