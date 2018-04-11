package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what should I do")
	}
}

func run() {
	newArgs := append([]string{"child"}, os.Args[2:]...)
	fmt.Printf("Running %v as %d\n", newArgs[0], os.Getpid())

	cmd := exec.Command("/proc/self/exe", newArgs...)

	// run it in UTS, PID and MNT namespaces
	// syscall.CLONE_NEWNS: MNT namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}

func child() {
        fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

        cmd := exec.Command(os.Args[2], os.Args[3:]...)

        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
