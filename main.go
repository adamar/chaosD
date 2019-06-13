package main

import (
	//"time"
	"fmt"
	_ "github.com/adamar/chaosd/plugins/inodepressure"
)

type Client struct {
	configFile string
	dbPath     string
}

func NewClient(configFile string) *Client {
	client := &Client{
		configFile: configFile,
		dbPath:     "/tmp/.db",
	}

	return client
}

func (c *Client) Run() {

	fmt.Println("Run chaosD")

}

func main() {

	chaosd := NewClient("/etc/chaod/config")
	chaosd.Run()

}
