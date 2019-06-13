

package plugin

type Plugin interface{}

var registeredPlugins = make([]Plugin, 0)

func RegisterPlugin(plugin Plugin) {
	registeredPlugins = append(registeredPlugins, plugin)
}

func RegisteredPlugins() []Plugin {
	return registeredPlugins
}
