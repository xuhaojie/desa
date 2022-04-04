package common

import "runtime"

type OsType int32

const (
	OS_UNKNOWN OsType = 0
	OS_WIN32   OsType = 1
	OS_LINUX   OsType = 2
	OS_DARWIN  OsType = 3
)

func (p OsType) String() string {
	switch p {
	case OS_WIN32:
		return "win32"
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
	case "win32":
		return OS_WIN32
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
