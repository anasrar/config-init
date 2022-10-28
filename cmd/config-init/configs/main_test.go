package configs

import (
	"errors"
	"os"
	"testing"
)

func TestList(t *testing.T) {
	for config, info := range List {
		for _, file := range info.Files {
			if _, err := os.Stat(file.PathSource); errors.Is(err, os.ErrNotExist) {
				t.Error("file not found", config, file.PathSource)
			}
		}
	}
}
