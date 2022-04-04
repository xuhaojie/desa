package setup

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"runtime"

	"autopard.com/desa/common"
)

func setupPipProxyLinux(mirror string) error {

	lines := []string{
		"[global]",
		"index-url=https://pypi.tuna.tsinghua.edu.cn/simple",
		"[install]",
		"trusted-host=https://pypi.tuna.tsinghua.edu.cn",
	}
	u, err := url.Parse(mirror)
	if err != nil {
		return err
	}

	lines[1] = fmt.Sprintf("index-url=%s", mirror)
	lines[3] = fmt.Sprintf("trusted-host=%s", u.Host)

	cfgSize := 0
	for _, line := range lines {
		cfgSize += len(line) + 1
	}
	cfg_data := make([]byte, 0, cfgSize)

	for _, line := range lines {
		cfg_data = append(cfg_data, line...)
		cfg_data = append(cfg_data, '\n')
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	targetPath := path.Join(userHomeDir, ".pip")
	targetFile := path.Join(targetPath, "pip.conf")
	exist, err := common.PathExists(targetPath)
	if err != nil {
		return err
	}
	if !exist {
		err := os.Mkdir(targetPath, 0664)
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(targetFile, []byte(cfg_data), os.FileMode(0644))
	return err
}

func SetupPipProxy(mirror string) error {
	switch runtime.GOOS {
	case "linux", "darwin":
		return setupPipProxyLinux(mirror)
	default:
		return errors.New("unsupported platform")
	}
}
