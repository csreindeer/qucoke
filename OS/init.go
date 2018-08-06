package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

//I am going to add input here, but for testing purposes I have not
//will create a for loop as an "engine"
	func main() {
		i := 1
		for i < 2 {
			root()
			output()
			i++
		}
	}

//will read qpuroot.py, via a sh work around
func root() {
	exec.Command("/bin/sh", "/home/mymainman/bigboy/qucoke/OS/rworkaround.sh").Run()
}

//will read qpu standard.py
func standard() {
	//run qpustandard
}
func output() {
	o, err := ioutil.ReadFile("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(o)
	fmt.Println(str)
}
