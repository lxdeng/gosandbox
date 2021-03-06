package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	case "grandchild":
		grandchild()
	default:
		panic("what should I do")
	}
}

func parent() {
	fmt.Println("enter parent: ....")

	// /proc/self/exe` which is a special file containing an in-memory image of the current executable.
	// i.e. re-run self and passing "child" as argument
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	// run it in UTS, PID and MNT namespaces
	// syscall.CLONE_NEWNS: MNT namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func child() {
	fmt.Println("enter child: ....")
	//cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd := exec.Command("/proc/self/exe", append([]string{"grandchild"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	time.Sleep(time.Minute * 3)

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func grandchild() {
	fmt.Println("enter grandchild: ....")
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	time.Sleep(time.Minute * 3)

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
