package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func setupPipProxyLinux() error {
	cfg := `
	[global]
	index-url = https://pypi.tuna.tsinghua.edu.cn/simple
	[install]
	trusted-host=mirrors.aliyun.com
	`
	userHomeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("User home directory: ", userHomeDir)
	}
	targetPath := path.Join(userHomeDir, ".pip")
	targetFile := path.Join(targetPath, "pip.conf")
	exist, err := PathExists(targetPath)
	if err != nil {
		return err
	}
	if !exist {
		err := os.Mkdir(targetPath, 0664)
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(targetFile, []byte(cfg), os.FileMode(0644))
	return err
}

func SetupPipProxy() error {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	switch runtime.GOOS {
	case "linux":
		return setupPipProxyLinux()
	}
	return nil
	// pip3 config list

}
