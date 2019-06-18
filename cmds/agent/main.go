package main

import (
	"fmt"
	"encoding/json"
	"github.com/adamar/chaosd/plugins"
	"github.com/adamar/chaosd/common"
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
		fmt.Println("Using default agent config")
	} else {
		fmt.Println("Reading agent config in from disk")
	}

	// get dbpath 
	c.dbPath = viper.GetString("dbPath")

}


func (c *Client) loadExperiments() {

	// Placeholder code until proper experiment loading is sorted
	js, _ := common.GetJson("https://chaosd.s3.amazonaws.com/example-experiment.json")
	config := make(map[string]string)
	err := json.Unmarshal(js, &config)
	if err != nil {
		fmt.Println("Config file loading error")
	}
	//config := make(map[string]string)
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
