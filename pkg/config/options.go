package config

import (
	"flag"
)

const (
	JSONOutput  = "json"
	TableOutput = "table"
)

type Options struct {
	Plugin     *string
	Mode       *string
	ScanPath   *string
	ScanParams *string
}

func ParseOptions() (*Options, error) {
	options := &Options{
		Plugin:     flag.String("plugin", "", "Path to the plugin to load"),
		Mode:       flag.String("mode", "runonce", "Mode to run the tool: runonce or server"),
		ScanPath:   flag.String("path", ".", "Path to scan"),
		ScanParams: flag.String("params", "", "Parameters to pass to the plugin"),
	}
	flag.Parse()
	return options, nil
}
