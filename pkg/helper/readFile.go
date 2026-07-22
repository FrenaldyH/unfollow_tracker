package helper

import (
	"os"

	"unfollow_tracker/pkg/logger"
)

func ReadFile(path string) ([]byte, error) {
	if byteValue, err := os.ReadFile(path); err != nil {
		logger.Log.Error("FILE not found",
			"error", err.Error(),
			"file", path,
		)
		return nil, err
	} else {
		return byteValue, nil
	}
}
