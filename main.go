package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Fuckoff")
	}

}
