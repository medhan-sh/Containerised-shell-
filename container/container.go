package container

import (
	"fmt"
	"os"
	"os/exec"
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
func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}
func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, "ERROR:", err)
	}
}
