package utils

import (
	"fmt"
	"runtime"
)

var (
	buildTime string
	gitHash   string
)

func Version() string {
	return fmt.Sprintf(`go version: %v
go os:      %v
go arch:    %v
build time: %v
git hash:   %v`,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		buildTime,
		gitHash,
	)
}
