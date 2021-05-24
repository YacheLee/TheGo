package main

import "os"

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	println(s)

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}

	//for true{
	//	println("HELLO")
	//}
}
