package main

import (
	"fmt"
	"github.com/adamar/chaosd/plugins"
	//_ "github.com/adamar/chaosd/plugins/inodepressure"
	_ "github.com/adamar/chaosd/plugins/networklatency"
	_ "github.com/adamar/chaosd/plugins/testplugin"
	"github.com/spf13/viper"
)

type Client struct {
	configFile       string
	dbPath           string
	installedPlugins map[string]plugins.Creator
}



func NewClient(configFile string) *Client {
	client := &Client{
		configFile:       configFile,
		installedPlugins: plugins.EnabledPlugins,
	}

	return client
}

func (c *Client) Run() {

	fmt.Println("Starting chaosD...")
	c.loadAgentConfig()
	c.loadExperiments()

}


func (c *Client) loadAgentConfig() {

	// name of config file (without extension)
	viper.SetConfigName(c.configFile)
	// path to look for the config file in
	viper.AddConfigPath(".")
	// set config type
	viper.SetConfigType("yaml")

	// set config defaults
        viper.SetDefault("dbPath", "/tmp/.db")

	// read in config
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	// get dbpath 
	c.dbPath = viper.GetString("dbPath")

}


func (c *Client) loadExperiments() {

	// Placeholder code until proper experiment loading is sorted
	config := make(map[string]string)
        config["Level"] = "5"
	config["Duration"] = "10"
        cons := plugins.EnabledPlugins["Testplugin"]
        experiment := cons(config)
        fmt.Printf("%+v\n", experiment)


}

func main() {

	chaosd := NewClient("example-config")
	chaosd.Run()

}
