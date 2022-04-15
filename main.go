package main

import (
	"flag"
	"fmt"
	"log"

	"autopard.com/desa/base"
	"autopard.com/desa/clean"
	"autopard.com/desa/download"
	"autopard.com/desa/setup"
)

var Cmd = &base.Command{
	UsageLine: "go",
	Long:      `Go is a tool for managing Go source code.`,
	// Commands initialized in package main
	Commands: []*base.Command{
		setup.Cmd,
		download.Cmd,
		clean.Cmd,
	},
}

func mainUsage() {
	//help.PrintUsage(os.Stderr, base.Go)
}

func main() {
	flag.Usage = base.Usage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		//base.Usage()
		return

	}

	if args[0] == "get" || args[0] == "help" {
		/*
			if !modload.WillBeEnabled() {
				// Replace module-aware get with GOPATH get if appropriate.
				*modget.CmdGet = *get.CmdGet
			}
		*/
	}
	/*
		cfg.CmdName = args[0] // for error messages
		if args[0] == "help" {
			help.Help(os.Stdout, args[1:])
			return
		}
	*/
	var found bool = false
	currentCmd := Cmd
	for _, cmd := range currentCmd.Commands {
		//fmt.Println(cmd.Name())
		//fmt.Println(args[1])

		if cmd.Name() != args[0] {
			continue
		} else {
			found = true
			if len(cmd.Commands) > 0 {
				currentCmd = cmd
				args = args[1:]
				if len(args) == 0 {

				}
				if args[0] == "help" {
					// Accept 'go mod help' and 'go mod help foo' for 'go help mod' and 'go help mod foo'.
					return
				}
			}
			cmd.Run(cmd, args[1:])
			break
		}

	}
	if !found {
		fmt.Println("unsurported command", args[0])
	}

}
