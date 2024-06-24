package scanx

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/ibreakthecloud/scanx/plugins"
)

type ScanX struct {
	plugins map[string]plugins.PluginInterface
}

func NewX() *ScanX {
	return &ScanX{
		plugins: make(map[string]plugins.PluginInterface),
	}
}

func (x *ScanX) LoadPlugin(name string) error {
	path, ok := plugins.PluginPathMap[name]
	if !ok {
		return fmt.Errorf("plugin %s not found", name)
	}

	p, err := plugin.Open(path)
	if err != nil {
		return err
	}

	symbol, err := p.Lookup("Plugin")
	if err != nil {
		return err
	}

	pluginInstance, ok := symbol.(*plugins.PluginInterface)
	if !ok {
		return fmt.Errorf("unexpected type from module symbol")
	}

	plug := *pluginInstance

	err = plug.Initialize()
	if err != nil {
		return err
	}

	x.plugins[plug.Name()] = plug
	return nil
}

func (x *ScanX) FileWalk(path string) {
	filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, pluginInstance := range x.plugins {
				pluginInstance.ProcessFile(filePath)
			}
		}
		return nil
	})
}
