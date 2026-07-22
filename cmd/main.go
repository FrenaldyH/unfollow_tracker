// cmd/main.go
package main

import (
	"unfollow_tracker/internal/gui"
	"unfollow_tracker/pkg/logger"
)

func main() {
	logger.Init()
	defer logger.Close()

	gui.Run()
}
