package download

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"autopard.com/desa/common"
)

type BuildType int32

const (
	BUILD_UNKNOWN BuildType = 0
	BUILD_STABLE  BuildType = 1
	BUILD_INSIDER BuildType = 2
)

func (p BuildType) String() string {
	switch p {
	case BUILD_STABLE:
		return "stable"
	case BUILD_INSIDER:
		return "insider"
	default:
		return "unknown"
	}
}

func genVscodeUrl(build BuildType, os common.OsType, arch common.ArchType, pkg common.PackageType) (string, error) {
	base := "https://code.visualstudio.com/sha/download"
	var result string
	switch os {
	case common.OS_WINDOWS:
		os_str := "win32"
		var arch_str string
		switch arch {
		case common.ARCH_X86:
			arch_str = ""
		case common.ARCH_AMD64:
			arch_str = "x64"
		case common.ARCH_ARM:
			return "", errors.New("arch arm not supported on win32 platform")
		case common.ARCH_ARM64:
			arch_str = "arm64"
		default:
			return "", errors.New("arch not supported on win32 platform")
		}

		switch pkg {
		case common.PACKAGE_EXE, common.PACKAGE_MSI, common.PACKAGE_UNKNOWN:
			if len(arch_str) > 0 {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s", base, build, os_str, arch_str)
			} else {
				result = fmt.Sprintf("%s?build=%s&os=%s", base, build, os_str)
			}

		case common.PACKAGE_ARCHIVE:
			if len(arch_str) > 0 {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s-%s", base, build, os_str, arch_str, "archive")
			} else {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s", base, build, os_str, "archive")
			}

		default:
			return "", errors.New("package type not supported on win32 platform")

		}
	case common.OS_LINUX:
		os_str := "linux"
		var arch_str string
		switch arch {
		case common.ARCH_AMD64:
			arch_str = "x64"
		case common.ARCH_ARM:
			arch_str = "armhf"
		case common.ARCH_ARM64:
			arch_str = "arm64"
		default:
			return "", errors.New("arch not supported on linux platform")
		}

		switch pkg {
		case common.PACKAGE_DEB:
			result = fmt.Sprintf("%s?build=%s&os=%s-%s-%s", base, build, os_str, "deb", arch_str)
		case common.PACKAGE_RPM:
			result = fmt.Sprintf("%s?build=%s&os=%s-%s-%s", base, build, os_str, "rpm", arch_str)
		case common.PACKAGE_ARCHIVE, common.PACKAGE_UNKNOWN:
			result = fmt.Sprintf("%s?build=%s&os=%s-%s", base, build, os, arch_str)
		}
	case common.OS_DARWIN:
		os_str := "darwin"
		var arch_str string
		switch arch {
		case common.ARCH_AMD64, common.ARCH_UNKNOWN:
			arch_str = ""
		case common.ARCH_ARM64:
			arch_str = "arm64"
		case common.ARCH_UNIVERSAL:
			arch_str = "universal"
		default:
			return "", errors.New("arch not supported on darwin platform")
		}
		if len(arch_str) > 0 {
			result = fmt.Sprintf("%s?build=%s&os=%s-%s", base, build, os_str, arch_str)
		} else {
			result = fmt.Sprintf("%s?build=%s&os=%s", base, build, os_str)
		}
	}

	return result, nil
}

func replaceVscodeDownloadUrl(url string, build BuildType, newbase string) string {
	// newbase = "https://vscode.cdn.azure.cn"
	//https: //vscode.cdn.azure.cn/stable/b4c1bd0a9b03c749ea011b06c6d2676c8091a70c/VSCodeUserSetup-x64-1.57.0.exe
	index := strings.Index(url, "/"+build.String()+"/")
	if index > 0 {
		return newbase + url[index:]
	} else {
		return url
	}
}

func DownloadVscode(buildType BuildType, osType common.OsType, archType common.ArchType, pkgType common.PackageType) error {

	url, err := genVscodeUrl(buildType, osType, archType, pkgType)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(url)

	url, err = common.GetFinalUrl(url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(url)

	targetUrl := replaceVscodeDownloadUrl(url, buildType, "https://vscode.cdn.azure.cn")
	fmt.Println("Get target url:", targetUrl)
	fileds := strings.Split(targetUrl, "/")
	file := fileds[len(fileds)-1]
	tmpDir := os.TempDir()

	common.DownloadFileProgress(targetUrl, path.Join(tmpDir, file))
	return nil
}
