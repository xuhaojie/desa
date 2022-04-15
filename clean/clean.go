package clean

import (
	"flag"
	"fmt"

	"autopard.com/desa/base"
)

type ProjectHandler interface {
	Detect(dir string)
	Clean()
}

type ProjectType int32

const (
	PROJECT_UNKNOWN      ProjectType = 0
	PROJECT_CARGO        ProjectType = 1
	PROJECT_GO           ProjectType = 2
	PROJECT_VISUALSTUDIO ProjectType = 3
)

type LanguageType int32

const (
	LANGUAGE_UNKNOWN LanguageType = 0
	LANGUAGE_RUST    LanguageType = 1
	LANGUAGE_GO      LanguageType = 2
	LANGUAGE_CPP     LanguageType = 3
)

type ProjectInfo struct {
	Dir  string
	Type ProjectType
}

var Cmd = &base.Command{
	UsageLine: "clean",
	Long:      `clean build cache.`,
	// Commands initialized in package main
	Run: Run,
}

func Run(cmd *base.Command, args []string) {
	cmdSet := flag.NewFlagSet("setup", flag.ContinueOnError)
	//var g = setupSet.Bool("g", false, "generate sample config")
	//var l = setupSet.Bool("l", false, "list targets")
	var p = cmdSet.String("p", "", "project path to clean")
	//var n = setupSet.String("n", "", "name of target to wake")
	//	setupSet.Parse(args)
	if len(args) < 1 {
		fmt.Printf("miss target")
		return
	}

	cmdSet.Parse(args[1:])
	target := args[0]
	fmt.Println("clean", target)
	switch target {
	case "rust":
		err := CleanRustProjects(*p)
		if err != nil {
			fmt.Println(err)
		}

	case "go":
		err := CleanGoProjects(*p)
		if err != nil {
			fmt.Println(err)
		}

	case "cpp":
		err := CleanCppProjects(*p)
		if err != nil {
			fmt.Println(err)
		}
	}
}
