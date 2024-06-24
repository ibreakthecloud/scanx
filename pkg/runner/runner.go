package runner

import (
	"fmt"

	"github.com/ibreakthecloud/scanx/pkg/scanx"
	log "github.com/sirupsen/logrus"
)

func RunOnce(x *scanx.ScanX, path string) {
	log.Infof("Running once on path: %s", path)
	x.FileWalk(path)
}

func ServerMode() {
	fmt.Println("Server mode not implemented yet")
}
