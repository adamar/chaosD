package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"syscall"
)

func enableLatency() {
	cmd := "tc"
	args := []string{"qdisc", "add", "dev", "eth0", "root", "netem", "delay", "97ms"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func disableLatency() {
	cmd := "tc"
	args := []string{"qdisc", "del", "dev", "eth0", "root", "netem"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func listInterfaces() {

	intf, _ := net.Interfaces()
	for _, e := range intf {
		fmt.Println(e)
	}
}

func createFile(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		file, _ := os.Create(path)
		//if isError(err) { return }
		defer file.Close()
	}
}

func getInodeStats(mountPoint string) {
	//https://linux.die.net/man/2/statvfs

	var stat syscall.Statfs_t

	syscall.Statfs(mountPoint, &stat)

	fmt.Println(stat.Files)
	fmt.Println(stat.Ffree)
	fmt.Println(stat.Files - stat.Ffree)

}

func getBlockStats(mountPoint string) {
	//https://linux.die.net/man/2/statvfs

	var stat syscall.Statfs_t

	syscall.Statfs(mountPoint, &stat)

	// Available blocks * size per block = available space in bytes
	availSpace := stat.Bavail * uint64(stat.Bsize)
	fmt.Println(availSpace)

	totalBlocks := stat.Blocks
	fmt.Println(totalBlocks)

}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func createEmptyFile(loc string) {

	os.OpenFile(loc, os.O_RDONLY|os.O_CREATE, 0666)

}
