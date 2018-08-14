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
		t := 2
		for i < t {
			initos()
			readoutput()
			i++
		}
	}
	
func initos() {
	exec.Command("/bin/sh", "OSrunner.sh").Run()
}
func readoutput() {
	o, err := ioutil.ReadFile("output.txt")

	if err != nil {
		fmt.Println(err)
	}
	str := string(o)
	fmt.Println(str)
}
