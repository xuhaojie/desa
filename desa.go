package main

import (
	"fmt"

	"autopard.com/desa/env"
)

func main() {
	// var builds = [...]string{"stable", "insider"}
	// var oss = [...]string{"win32", "linux", "darwin"}
	// var archs = [...]string{"x64", "universal", "arm64"}
	// var formats = [...]string{"user", "archive", "deb", "rpm"}
	//	uri := genVscodeUrl("insider", "linux", "x64", "rpm")

	/*
		build := "stable"
		os := "linux"
		//arch := "arm64"
		arch := "x64"
		format := "deb"
		err := downloader.DownloadVscode(build, os, arch, format)
		if err != nil {
			fmt.Println(err)
		}
	*/

	err := env.SetupPipProxy()
	if err != nil {
		fmt.Println(err)
	}

	err = env.SetupGolangProxy()
	if err != nil {
		fmt.Println(err)
	}

	err = env.SetupCargoProxy()
	if err != nil {
		fmt.Println(err)
	}

}
