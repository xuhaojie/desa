package main

import (
	"fmt"

	"autopard.com/desa/common"
	"autopard.com/desa/download"
	"autopard.com/desa/setup"
)

func main() {
	// var builds = [...]string{"stable", "insider"}
	// var oss = [...]string{"win32", "linux", "darwin"}
	// var archs = [...]string{"x64", "universal", "arm64"}
	// var formats = [...]string{"user", "archive", "deb", "rpm"}
	//	uri := genVscodeUrl("insider", "linux", "x64", "rpm")
	flag_download_vscode := true
	flag_setup_pip := false
	flag_setup_go := false
	flag_setup_cargo := false
	flag_setup_git := false
	if flag_download_vscode {
		build := download.BUILD_STABLE

		//os := common.GetOsType()
		//os := common.OS_WIN32
		os := common.OS_LINUX
		//os := common.OS_DARWIN

		arch := common.GetArchType()
		//arch := common.ARCH_X86
		//arch := common.ARCH_AMD64
		//arch := common.ARCH_ARM
		//arch := common.ARCH_ARM64
		//arch := common.ARCH_UNIVERSAL

		pkg := download.PACKAGE_UNKNOWN
		//pkg := download.PACKAGE_EXE
		//pkg := download.PACKAGE_MSI
		//pkg := download.PACKAGE_DEB
		//pkg := download.PACKAGE_RPM
		//pkg := download.PACKAGE_ARCHIVE

		err := download.DownloadVscode(build, os, arch, pkg)
		if err != nil {
			fmt.Println(err)
		}
	}

	if flag_setup_pip {
		err := setup.SetupPipProxy()
		if err != nil {
			fmt.Println(err)
		}
	}

	if flag_setup_go {
		err := setup.SetupGolangProxy()
		if err != nil {
			fmt.Println(err)
		}
	}

	if flag_setup_cargo {
		err := setup.SetupCargoProxy()
		if err != nil {
			fmt.Println(err)
		}
	}

	if flag_setup_git {
		err := setup.SetupGit()
		if err != nil {
			fmt.Println(err)
		}
	}

}
