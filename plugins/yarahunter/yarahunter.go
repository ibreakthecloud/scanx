// plugins/malware/malware.go
package main

import (
	"fmt"

	"github.com/deepfence/YaraHunter/constants"
	"github.com/deepfence/YaraHunter/pkg/config"
	"github.com/deepfence/YaraHunter/pkg/scan"
	"github.com/deepfence/YaraHunter/pkg/yararules"
	"github.com/ibreakthecloud/scanx/plugins"
)

type MalwareScanner struct {
	Config    config.Config
	RulesPath string
	ScanParam ScanParam
	Scanner   *scan.Scanner
}

func (m *MalwareScanner) Name() string {
	return "MalwareScanner"
}

func (m *MalwareScanner) Initialize() error {
	yaraRules := yararules.New(m.RulesPath)
	// todo: remove hardcoding false
	err := yaraRules.Compile(constants.Filescan, false)
	if err != nil {
		return err
	}

	yaraScanner, err := yaraRules.NewScanner()
	if err != nil {
		return err
	}

	m.Scanner = scan.New(&config.Options{}, &config.Config{}, yaraScanner, "")

	fmt.Println("MalwareScanner initialized")
	return nil
}

func (m *MalwareScanner) ProcessFile(path string) error {
	fmt.Printf("Scanning file for malware: %s\n", path)
	return nil
}

var Plugin plugins.PluginInterface = &MalwareScanner{}
