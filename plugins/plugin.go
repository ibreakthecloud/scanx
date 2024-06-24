// plugin_interface.go
package plugins

type PluginInterface interface {
	Name() string
	Initialize() error
	ProcessFile(path string) error
}
