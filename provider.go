package main

import (
	"fmt"
	//"time"
)

type failure interface {
	start() string
	end() string
}

type latencySpike struct {
	networkIntface string
	level          string
}

func (f latencySpike) start() string {
	return "started latency spike"
}

func (f latencySpike) end() string {
	return "end latency spike"
}

type diskIOSpike struct {
	volume string
	level  string
}

func (f diskIOSpike) start() string {
	return "started disk io spike"
}

func (f diskIOSpike) end() string {
	return "end disk io spike"
}

func runTest(f failure) {
	fmt.Println(f)
	f.start()
	f.end()
}
