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
		t := 3
		for i < t {
			initos()
			readoutput()
			i++
		}
	}

//will read qpuroot.py, via a sh work around
func initos() {
	exec.Command("/bin/sh", "OSrunner.sh").Run()
}

//will read qpu standard.py
func standard() {
	//run qpustandard
}
func readoutput() {
	o, err := ioutil.ReadFile("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(o)
	fmt.Println(str)
}
