package setup

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

func SetupNpmProxyLinux(mirror string) error {
	cmds := []Cmd{
		{cmd: "npm", params: []string{"config", "set", "registry", mirror}},
	}

	out, err := executeCmds(cmds)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}

func SetupNpmProxy(mirror string) error {
	switch runtime.GOOS {
	case "linux", "darwin":
		return SetupNpmProxyLinux(mirror)
	default:
		return errors.New("unsupported platform")
	}
	return nil
	// pip3 config list

}
