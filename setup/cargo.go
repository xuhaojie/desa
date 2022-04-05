package setup

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"autopard.com/desa/common"
)

func SetupCargoProxy(mirror string) error {
	lines := []string{
		"[source.crates-io]",
		"registry = \"https://github.com/rust-lang/crates.io-index\"",
		"replace-with = 'mirror'",
		"[source.mirror]",
		"registry = \"https://mirrors.ustc.edu.cn/crates.io-index\"",
	}

	lines[4] = fmt.Sprintf("registry = \"%s\"", mirror)

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

	targetPath := path.Join(userHomeDir, ".cargo")
	targetFile := path.Join(targetPath, "config")
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
