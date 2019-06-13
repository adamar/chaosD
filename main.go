package main

import (
//"time"
)

func main() {

	//d := diskIOSpike{volume: "/dev/sda1", level: "medium"}
	//l := latencySpike{networkIntface: "eth0", level: "medium"}

	//runTest(d)
	//runTest(l)

	i := inodeUse{mountPoint: "/", level: 100, duration: 20}
	runTest(i)

}
