package setup

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"autopard.com/desa/common"
)

func setupCargoProxyLinux() error {
	cfg := `
	[source.crates-io]
	registry = "https://github.com/rust-lang/crates.io-index"
	# 指定镜像
	replace-with = 'tuna' # 如：tuna、sjtu、ustc，或者 rustcc
	
	# 注：以下源配置一个即可，无需全部
	
	# 中国科学技术大学
	[source.ustc]
	registry = "https://mirrors.ustc.edu.cn/crates.io-index"
	# >>> 或者 <<<
	#registry = "git://mirrors.ustc.edu.cn/crates.io-index"
	
	# 上海交通大学
	[source.sjtu]
	registry = "https://mirrors.sjtug.sjtu.edu.cn/git/crates.io-index/"
	
	# 清华大学
	[source.tuna]
	registry = "https://mirrors.tuna.tsinghua.edu.cn/git/crates.io-index.git"
	
	# rustcc社区
	[source.rustcc]
	registry = "https://code.aliyun.com/rustcc/crates.io-index.git"
	`
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

	err = ioutil.WriteFile(targetFile, []byte(cfg), os.FileMode(0644))
	return err
}

func SetupCargoProxy() error {
	switch runtime.GOOS {
	case "linux", "darwin":
		return setupCargoProxyLinux()
	default:
		return errors.New("unsupported platform")
	}
	return nil
	// pip3 config list

}
