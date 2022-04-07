package setup

import (
	"flag"
	"fmt"

	"autopard.com/desa/base"
)

var Cmd = &base.Command{
	UsageLine: "setup",
	Long:      `setup develop environment.`,
	// Commands initialized in package main
	Run: Run,
}

func Run(cmd *base.Command, args []string) {
	fmt.Println("main")
	setupSet := flag.NewFlagSet("setup", flag.ContinueOnError)
	//var g = setupSet.Bool("g", false, "generate sample config")
	//var l = setupSet.Bool("l", false, "list targets")
	var m = setupSet.String("m", "tuna", "mirror name")

	setupSet.Parse(args[1:])
	target := args[0]
	switch target {
	case "apt":
		var mirrorName string = "tuna"
		if m != nil {
			mirrorName = *m
		}
		var mirror string
		switch mirrorName {
		case "tuna":
			mirror = "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/"
		case "163":
			mirror = "http://mirrors.163.com/ubuntu/"
		case "aliyun":
			mirror = "https://mirrors.aliyun.com/ubuntu/"
		default:
			fmt.Println("unknown mirror")
			return
		}

		err := SetupAptSource(mirror)
		if err != nil {
			fmt.Println(err)
		}
	case "cargo":
		var mirrorName string = "tuna"
		if m != nil {
			mirrorName = *m
		}
		var mirror string
		switch mirrorName {
		case "utsc":
			mirror = "https://mirrors.ustc.edu.cn/crates.io-index"
		case "tuna":
			mirror = "https://mirrors.tuna.tsinghua.edu.cn/git/crates.io-index.git"
		case "sjtu":
			mirror = "https://mirrors.sjtug.sjtu.edu.cn/git/crates.io-index/"
		case "rustcc":
			mirror = "https://code.aliyun.com/rustcc/crates.io-index.git"
		default:
			fmt.Println("unknown mirror")
			return
		}

		err := SetupCargoProxy(mirror)
		if err != nil {
			fmt.Println(err)
		}
	case "git":
		err := SetupGit()
		if err != nil {
			fmt.Println(err)
		}
	case "go":
		var mirrorName string = "taobao"
		if m != nil {
			mirrorName = *m
		}
		var mirror string
		switch mirrorName {
		case "goproxy.io":
			mirror = "https://goproxy.io/zh/"
		case "goproxy.cn":
			mirror = "https://goproxy.cn"
		case "aliyun":
			mirror = "https://mirrors.aliyun.com/goproxy/"
		default:
			fmt.Println("unknown mirror")
			return
		}
		err := SetupGolangProxy(mirror)
		if err != nil {
			fmt.Println(err)
		}
	case "npm":
		var mirrorName string = "taobao"
		if m != nil {
			mirrorName = *m
		}
		var mirror string
		switch mirrorName {
		case "origin":
			mirror = "https://registry.npmjs.org/"
		case "taobao":
			mirror = "https://registry.npm.taobao.org"
		default:
			fmt.Println("unknown mirror")
			return
		}

		err := SetupNpmProxy(mirror)
		if err != nil {
			fmt.Println(err)
		}
	case "pip":
		var mirrorName string = "tuna"
		if m != nil {
			mirrorName = *m
		}
		var mirror string
		switch mirrorName {
		case "tuna":
			mirror = "https://pypi.tuna.tsinghua.edu.cn/simple"
		case "163":
			mirror = "https://mirrors.163.com/pypi/simple"
		case "aliyun":
			mirror = "http://mirrors.aliyun.com/pypi/simple"
		default:
			fmt.Println("unknown mirror")
			return
		}
		err := SetupPipProxy(mirror)
		if err != nil {
			fmt.Println(err)
		}

	}
}
