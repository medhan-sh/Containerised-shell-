package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func GetContainerPath() (string, string) {
	exe, _ := os.Executable()
	baseDir := filepath.Dir(exe)
	rootfs := filepath.Join(baseDir, "rootfs")
	proc := filepath.Join(rootfs, "proc")
	return rootfs, proc
}
func runContainer() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	if err := cmd.Run(); err != nil {
		fmt.Fprint(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}
func child() {
	rootfs, proc := GetContainerPath()
	if _, err := os.Stat(rootfs); os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "rootfs not found — run: make setup")
		os.Exit(1)
	}
	if err := syscall.Sethostname([]byte("Custom Shell")); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR settting Hostname :%v\n", err)
	}
	if err := syscall.Mount("proc", proc, "proc", 0, ""); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR setting up Mount :%v\n", err)
	}
	if err := syscall.Chroot(rootfs); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR Changing root :%v\n", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR changing directory:%v\n", err)
	}
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR:%v\n", err)
	}
}
