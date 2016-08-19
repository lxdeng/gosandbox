package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	//"time"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("wat should I do")
	}
}

func parent() {
	// /proc/self/exe` which is a special file containing an in-memory image of the current executable.
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

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

	must(os.MkdirAll("rootfs/oldrootfs", 0700))
	must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))

	// tell the OS to move the current directory at `/` to `rootfs/oldrootfs` , and to swap the new rootfs directory to `/` . After the `pivotroot` call is complete, the / directory in the container will refer to the rootfs. (The bind mount call is needed to satisfy some requirements of the `pivotroot` command — the OS requires that `pivotroot` be used to swap two filesystems that are not part of the same tree, which bind mounting the rootfs to itself achieves. Yes, it’s pretty silly).

	must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	must(os.Chdir("/"))

	walkDir("/oldrootfs/tmp")

	ls()

	/*
		cmd := exec.Command(os.Args[2], os.Args[3:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		//time.Sleep(time.Minute * 3)

		if err := cmd.Run(); err != nil {
			fmt.Println("ERROR", err)
			os.Exit(1)
		}
	*/
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func ls() {
	path, err := exec.LookPath("/oldrootfs/bin/ls")
	if err != nil {
		log.Fatal("ls command not found")
	}

	fmt.Printf("ls is available at %s\n", path)

	must(os.Chdir("/oldrootfs/bin"))

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("os.Getwd failed: ", err)
	}
	fmt.Printf("cwd=%s\n", cwd)

	cmd := exec.Command(path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err) // why cmd.Run() cannot find the executable ls?
		os.Exit(1)
	}

}

func walkDir(dir string) {
	err := filepath.Walk(dir, visit)
	if err != nil {
		log.Fatal(err)
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	fmt.Printf("%s isDir=%v\n", path, info.IsDir())
	return nil
}
