package main

import "os"

var cwd string = "adsada"

func init() {
	cwd, err := os.Getwd()
	if err != nil {

	}
	println("OUT", cwd)
}

func main() {
	println(cwd)
}
