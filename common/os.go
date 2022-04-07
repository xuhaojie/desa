package common

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

type OsType int32

const (
	OS_UNKNOWN OsType = 0
	OS_WINDOWS OsType = 1
	OS_LINUX   OsType = 2
	OS_DARWIN  OsType = 3
)

func (p OsType) String() string {
	switch p {
	case OS_WINDOWS:
		return "windows"
	case OS_LINUX:
		return "linux"
	case OS_DARWIN:
		return "darwin"
	default:
		return "unknown"
	}
}

func GetOsType() OsType {
	switch runtime.GOOS {
	case "linux":
		return OS_LINUX
	case "windows":
		return OS_WINDOWS
	case "darwin":
		return OS_DARWIN
	}
	return OS_UNKNOWN
}

type ArchType int32

const (
	ARCH_UNKNOWN   ArchType = 0
	ARCH_X86       ArchType = 1
	ARCH_AMD64     ArchType = 2
	ARCH_ARM       ArchType = 3
	ARCH_ARM64     ArchType = 4
	ARCH_UNIVERSAL ArchType = 0xff
)

func GetArchType() ArchType {
	switch runtime.GOARCH {
	case "x86":
		return ARCH_X86
	case "amd64":
		return ARCH_AMD64
	case "arm":
		return ARCH_ARM
	case "arm64":
		return ARCH_ARM64
	}
	return ARCH_UNKNOWN
}

type OsInfo struct {
	Type     string
	Name     string
	Version  string
	CodeName string
}

func (i OsInfo) String() string {
	return fmt.Sprintf("type: %s\nname: %s\nversion: %s\ncodename: %s\n", i.Type, i.Name, i.Version, i.CodeName)
}

func GetOsVersionInfo() OsInfo {
	var osInfo OsInfo
	switch runtime.GOOS {
	case "linux":
		command := exec.Command("lsb_release", "-a")
		out, err := command.CombinedOutput()
		if err != nil {
			log.Fatalf("command.Run() failed with %s\n", err)
		}

		var output = string(out)
		fileds := strings.Split(output, "\n")
		var infoMap = make(map[string]string)

		for _, filed := range fileds {
			if strings.Index(filed, ":") < 0 {
				continue
			}
			v := strings.Split(filed, ":")

			key := v[0]
			value := strings.TrimLeft(v[1], " \t")
			//value = strings.TrimRight(value, " \t")
			infoMap[key] = value
		}
		//fmt.Println(infoMap)

		osInfo.Type = "linux"
		osInfo.Name = infoMap["Distributor ID"]
		osInfo.Version = infoMap["Release"]
		osInfo.CodeName = infoMap["Codename"]
	case "darwin":
		osInfo.Type = "darwin"
	case "win32":
		osInfo.Type = "windows"

	}

	return osInfo
}
