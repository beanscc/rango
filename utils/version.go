package utils

import (
	"fmt"
	"runtime"
)

var (
	buildTime string
	gitBranch string
	gitHash   string
)

func Version() string {
	return fmt.Sprintf(`go version: %v
go os:      %v
go arch:    %v
build time: %v
git branch: %v
git hash:   %v`,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		buildTime,
		gitBranch,
		gitHash,
	)
}
