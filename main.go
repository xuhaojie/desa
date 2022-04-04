package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"autopard.com/desa/common"
	"autopard.com/desa/download"
	"autopard.com/desa/setup"
)

var (
	required string

	downloadCmd = flag.NewFlagSet("download", flag.ContinueOnError)
	setupCmd    = flag.NewFlagSet("setup", flag.ContinueOnError)
)

var subCommands = map[string]*flag.FlagSet{
	downloadCmd.Name(): downloadCmd,
	setupCmd.Name():    setupCmd,
}

func downloadCmdHandler(args []string) {
	downloadSet := flag.NewFlagSet("setup", flag.ContinueOnError)
	//var g = setupSet.Bool("g", false, "generate sample config")
	//var l = setupSet.Bool("l", false, "list targets")
	//var m = setupSet.String("m", "", "mac address of target to wake")
	//var n = setupSet.String("n", "", "name of target to wake")
	//	setupSet.Parse(args)
	downloadSet.Parse(args[1:])
	target := args[0]
	fmt.Println("download", target)
	switch target {
	case "vscode":
		build := download.BUILD_STABLE

		os := common.GetOsType()
		//os := common.OS_WIN32
		//os := common.OS_LINUX
		//os := common.OS_DARWIN

		arch := common.GetArchType()
		//arch := common.ARCH_X86
		//arch := common.ARCH_AMD64
		//arch := common.ARCH_ARM
		//arch := common.ARCH_ARM64
		//arch := common.ARCH_UNIVERSAL

		pkg := download.PACKAGE_UNKNOWN

		if os == common.OS_LINUX && pkg == download.PACKAGE_UNKNOWN {
			osInfo := common.GetOsVersionInfo()
			//fmt.Println(osInfo)
			switch strings.ToLower(osInfo.Name) {
			case "ubuntu", "debian":
				pkg = download.PACKAGE_DEB
			case "centos", "fordera":
				pkg = download.PACKAGE_RPM
			default:
				pkg = download.PACKAGE_ARCHIVE
			}
		}

		err := download.DownloadVscode(build, os, arch, pkg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func setupCmdHandler(args []string) {
	setupSet := flag.NewFlagSet("setup", flag.ContinueOnError)
	//var g = setupSet.Bool("g", false, "generate sample config")
	//var l = setupSet.Bool("l", false, "list targets")
	//var m = setupSet.String("m", "", "mac address of target to wake")
	//var n = setupSet.String("n", "", "name of target to wake")
	setupSet.Parse(args[1:])
	target := args[0]
	switch target {
	case "cargo":
		err := setup.SetupCargoProxy()
		if err != nil {
			fmt.Println(err)
		}
	case "git":
		err := setup.SetupGit()
		if err != nil {
			fmt.Println(err)
		}
	case "go":
		err := setup.SetupGolangProxy()
		if err != nil {
			fmt.Println(err)
		}
	case "pip":
		err := setup.SetupPipProxy()
		if err != nil {
			fmt.Println(err)
		}

	}
}

func main() {

	//	configFile := filepath.Join(configdir.LocalConfig(), "gwaker.cfg")
	if len(os.Args) < 3 {
		fmt.Println("expected 'setup' or 'download' subcommands")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "download":
		downloadCmdHandler(os.Args[2:])
	case "setup":
		setupCmdHandler(os.Args[2:])
	}
	/*
		if required == "" {
			fmt.Println("-required is required for all commands")
		}
	*/
}
