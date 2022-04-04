package main

import (
	"fmt"

	"autopard.com/desa/downloader"
	"autopard.com/desa/env"
)

func main() {
	// var builds = [...]string{"stable", "insider"}
	// var oss = [...]string{"win32", "linux", "darwin"}
	// var archs = [...]string{"x64", "universal", "arm64"}
	// var formats = [...]string{"user", "archive", "deb", "rpm"}
	//	uri := genVscodeUrl("insider", "linux", "x64", "rpm")
	flag_download_vscode := false
	flag_setupenv_pip := true
	flag_setupenv_go := true
	flag_setupenv_cargo := true
	if flag_download_vscode {
		build := "stable"
		os := "darwin"
		//arch := "arm64"
		arch := ""
		format := ""
		err := downloader.DownloadVscode(build, os, arch, format)
		if err != nil {
			fmt.Println(err)
		}
	}
	if flag_setupenv_pip {
		err := env.SetupPipProxy()
		if err != nil {
			fmt.Println(err)
		}

	}
	if flag_setupenv_go {
		err := env.SetupGolangProxy()
		if err != nil {
			fmt.Println(err)
		}
	}

	if flag_setupenv_cargo {
		err := env.SetupCargoProxy()
		if err != nil {
			fmt.Println(err)
		}
	}

}
