package testplugin

import (
	"github.com/adamar/chaosd/plugins"
)

type Testplugin struct {
	Level string
	Duration string
}

func init() {
	plugins.Add("Testplugin", func(config map[string]string) plugins.Plugin {
		return &Testplugin{Level: config["Level"], Duration: config["Duration"]}
	})
}

func (t *Testplugin) Description() string {
	return "Example plugin"
}

func (t *Testplugin) Start() string {
	return "Starting"
}

func (t *Testplugin) End() string {
	return "Ending"
}
