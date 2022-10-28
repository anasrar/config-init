package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/anasrar/config-init/cmd/config-init/configs"
	"github.com/fatih/color"
	"github.com/samber/lo"
)

//go:embed VERSION.txt
var VERSION string

func main() {
	args := lo.Uniq(os.Args[1:])
	if len(args) == 0 {
		yellow := color.New(color.FgYellow).Add(color.Bold)
		green := color.New(color.FgGreen).Add(color.Bold)

		fmt.Print("config-init " + VERSION + "\n" + yellow.Sprint("USAGE") + "\n    config-init [...CONFIGS]\n\n" + yellow.Sprint("CONFIGS") + "\n")

		longestConfigName := longestString(lo.Keys(configs.List))
		for configName, info := range configs.List {
			gap := longestConfigName - len(configName) + 3
			files := lo.Map(info.Files, func(item configs.File, _ int) string {
				return item.PathTarget
			})
			fmt.Println(green.Sprint("    "+configName) + strings.Repeat(" ", gap) + strings.Join(files, ", "))
		}
	} else {
		red := color.New(color.FgRed).Add(color.Bold)
		green := color.New(color.FgGreen).Add(color.Bold)

		longestConfigName := longestString(args)
		for _, configName := range args {
			info, exist := configs.List[configName]
			gap := longestConfigName - len(configName) + 3

			if !exist {
				fmt.Println(red.Sprint(configName) + strings.Repeat(" ", gap) + "config not found")
			} else {
				files := lo.Map(info.Files, func(item configs.File, _ int) string {
					return item.PathTarget
				})

				for _, file := range info.Files {
					createFile(file.PathTarget, file.PathSource)
				}

				fmt.Println(green.Sprint(configName) + strings.Repeat(" ", gap) + "create " + strings.Join(files, ", "))
			}
		}
	}
}

func longestString(items []string) int {
	if len(items) == 0 {
		return 0
	}
	if len(items) == 1 {
		return len(items[0])
	}
	return lo.Reduce(items[1:], func(agg int, item string, _ int) int {
		return int(math.Max(float64(agg), float64(len(item))))
	}, len(items[0]))
}

func createFile(target string, source string) {
	if err := os.WriteFile(target, configs.GetConfig(source), 0644); err != nil {
		log.Fatal(err)
	}
}
