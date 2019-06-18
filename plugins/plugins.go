package plugins

var EnabledPlugins = map[string]Creator{}

type Creator func(config map[string]string) Plugin

func Add(name string, creator Creator) {
	EnabledPlugins[name] = creator
}

type Plugin interface {
	// Description returns a one-sentence description on the Input
	Description() string
	Start() string
	End() string
}
