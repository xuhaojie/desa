package setup

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"autopard.com/desa/common"
)

func reverseSearch(s string, target byte) int {
	var t int
	for t = len(s) - 1; t >= 0; t-- {
		if s[t] == '/' {
			return t
		}
	}
	return -1
}

type MirrorName int32

const (
	MIRROR_163    MirrorName = 1
	MIRROR_TUNA   MirrorName = 2
	MIRROR_ALIYUN MirrorName = 3
)

func (m MirrorName) String() string {
	switch m {
	case MIRROR_163:
		return "http://mirrors.163.com/ubuntu/"
	case MIRROR_TUNA:
		return "https://mirrors.tuna.tsinghua.edu.cn/ubuntu/"
	case MIRROR_ALIYUN:
		return "https://mirrors.aliyun.com/ubuntu/"
	default:
		return "unknown"
	}
}

func setupAptSourceLinux() error {
	mirrorName := MIRROR_ALIYUN
	mirrorUrl := mirrorName.String()
	fmt.Println(mirrorUrl)

	targetPath := "/etc/apt/sources.list"
	backupPath := targetPath + ".backup"

	readFile, err := os.Open(targetPath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	fileSize := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fileSize += len(line)

		start := strings.Index(line, "deb http")
		if start > 0 {
			end := reverseSearch(line, '/')
			if end > start {
				target := line[start+4 : end+1]
				//fmt.Println(target)
				line = strings.Replace(line, target, mirrorUrl, -1)
			}
		} else {
			start := strings.Index(line, "deb-src http")
			if start > 0 {
				end := reverseSearch(line, '/')
				if end > start {
					target := line[start+8 : end+1]
					//fmt.Println(target)
					line = strings.Replace(line, target, mirrorUrl, -1)

				}
			}
		}
		lines = append(lines, line)
	}

	readFile.Close()

	cfg_data := make([]byte, 0, fileSize)

	for _, line := range lines {
		cfg_data = append(cfg_data, line...)
		cfg_data = append(cfg_data, '\n')
	}

	exist, err := common.PathExists(targetPath)
	if !exist {
		return errors.New("config file does not exist")
	}

	err = os.Rename(targetPath, backupPath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(targetPath, cfg_data, os.FileMode(0644))
	return err
}

func SetupAptSource() error {
	switch runtime.GOOS {
	case "linux":
		return setupAptSourceLinux()
	default:
		return errors.New("unsupported platform")
	}
	return nil
	// pip3 config list

}
