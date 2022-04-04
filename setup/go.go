package setup

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

func setupGolangLinux() error {
	//$ go env -w GO111MODULE=on
	//$ go env -w GOPROXY=https://goproxy.cn,directory
	cmds := []Cmd{
		{cmd: "go", params: []string{"env", "-w", "GO111MODULE=on"}},
		{cmd: "go", params: []string{"env", "-w", "GOPROXY=https://goproxy.cn,direct"}},
	}
	out, err := executeCmds(cmds)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}

func SetupGolangProxy() error {
	switch runtime.GOOS {
	case "linux", "darwin":
		return setupGolangLinux()
	default:
		return errors.New("unsupported platform")
	}
	return nil
	// pip3 config list

}
