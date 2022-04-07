package download

import (
	"flag"
	"fmt"
	"runtime"
	"strings"

	"autopard.com/desa/base"
	"autopard.com/desa/common"
)

var Cmd = &base.Command{
	UsageLine: "download",
	Long:      `download software.`,
	// Commands initialized in package main
	Run: Run,
}

func Run(cmd *base.Command, args []string) {

	downloadSet := flag.NewFlagSet("setup", flag.ContinueOnError)
	//var g = setupSet.Bool("g", false, "generate sample config")
	//var l = setupSet.Bool("l", false, "list targets")
	//var m = setupSet.String("m", "", "mac address of target to wake")
	//var n = setupSet.String("n", "", "name of target to wake")
	//	setupSet.Parse(args)
	if len(args) < 1 {
		return
	}
	downloadSet.Parse(args[1:])
	target := args[0]
	fmt.Println("download", target)
	build := BUILD_STABLE

	os := common.GetOsType()
	//os := common.OS_WINDOWS
	//os := common.OS_LINUX
	//os := common.OS_DARWIN

	arch := common.GetArchType()
	//arch := common.ARCH_X86
	//arch := common.ARCH_AMD64
	//arch := common.ARCH_ARM
	//arch := common.ARCH_ARM64
	//arch := common.ARCH_UNIVERSAL

	pkg := common.PACKAGE_UNKNOWN

	switch target {
	case "vscode":
		if os == common.OS_LINUX && pkg == common.PACKAGE_UNKNOWN {
			osInfo := common.GetOsVersionInfo()
			//fmt.Println(osInfo)
			switch strings.ToLower(osInfo.Name) {
			case "ubuntu", "debian":
				pkg = common.PACKAGE_DEB
			case "centos", "fordera":
				pkg = common.PACKAGE_RPM
			default:
				pkg = common.PACKAGE_ARCHIVE
			}
		}
		err := DownloadVscode(build, os, arch, pkg)
		if err != nil {
			fmt.Println(err)
		}
	case "nomachine":
		switch runtime.GOOS {
		case "linux":
			if pkg == common.PACKAGE_UNKNOWN {
				osInfo := common.GetOsVersionInfo()
				//fmt.Println(osInfo)
				switch strings.ToLower(osInfo.Name) {
				case "ubuntu", "debian":
					pkg = common.PACKAGE_DEB
				case "centos", "fordera":
					pkg = common.PACKAGE_RPM
				default:
					pkg = common.PACKAGE_ARCHIVE
				}
			}
		case "windows":
			if pkg == common.PACKAGE_UNKNOWN {
				pkg = common.PACKAGE_EXE
			}
		}
		err := DownloadNomachine(build, os, arch, pkg)
		if err != nil {
			fmt.Println(err)
		}
	case "vmware":
		err := DownloadVmware(build, os, arch, pkg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
