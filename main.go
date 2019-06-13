package main

import (
	//"time"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/adamar/chaosd/plugins/inodepressure"
)

type Client struct {
	configFile string
	dbPath     string
}

func NewClient(configFile string) *Client {
	client := &Client{
		configFile: configFile,
	}

	return client
}

func (c *Client) Run() {

	fmt.Println("Run chaosD")
	c.loadConfig()

}

func (c *Client) loadConfig() {

        viper.SetConfigName("example-config") // name of config file (without extension)
        viper.AddConfigPath(".")   // path to look for the config file in

        err := viper.ReadInConfig()
        if err != nil {
	    fmt.Println(err.Error())
        }

	viper.SetConfigType("yaml")
        viper.SetDefault("dbPath", "/tmp/.db")

	c.dbPath = viper.GetString("dbPath")


}

func main() {

	chaosd := NewClient("config.yml")
	chaosd.Run()

}
