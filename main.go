package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "run":
		runContainer()
	case "child":
		child()
	default:
		panic("Fuckoff")
	}

}
