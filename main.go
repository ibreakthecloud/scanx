// main.go
package main

import (
	"github.com/ibreakthecloud/scanx/pkg/config"
	"github.com/ibreakthecloud/scanx/pkg/runner"
	"github.com/ibreakthecloud/scanx/pkg/scanx"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.SetLogger()
	opts, err := config.ParseOptions()
	if err != nil {
		log.Panicf("main: failed to parse options: %v", err)
	}

	if *opts.Plugin == "" {
		log.Fatal("Please specify a plugin using the -plugin flag")
		return
	}

	x := scanx.NewX()
	err = x.LoadPlugin(*opts.Plugin)
	if err != nil {
		log.Fatal(err)
		return
	}

	switch *opts.Mode {
	case "runonce":
		runner.RunOnce(x, *opts.ScanPath)
	case "server":
		runner.ServerMode()
	default:
		log.Fatal("Unknown mode. Use 'runonce' or 'server'")
	}
}
