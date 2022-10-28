package configs

import (
	"embed"
	"log"
)

//go:embed raws/*
var raws embed.FS

func GetConfig(path string) []byte {
	result, err := raws.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

type File struct {
	PathTarget string
	PathSource string
}

type Info struct {
	Files []File
}

var List = map[string]Info{
	"editorconfig": {
		Files: []File{{
			PathTarget: ".editorconfig",
			PathSource: "raws/.editorconfig",
		}},
	},
	"prettier": {
		Files: []File{{
			PathTarget: ".prettierrc.yml",
			PathSource: "raws/.prettierrc.yml",
		}},
	},
}
