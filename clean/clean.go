package clean

import (
	"flag"
	"fmt"

	"autopard.com/desa/base"
)

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
