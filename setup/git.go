package setup

import (
	"fmt"
	"log"
	"runtime"
)

func setupGitLinux() error {
	cmds := []Cmd{
		{cmd: "git", params: []string{"config", "--global", "user.name", "xuhaojie"}},
		{cmd: "git", params: []string{"config", "--global", "user.email", "xuhaojie@hotmail.com"}},
	}

	out, err := executeCmds(cmds)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}

func SetupGit() error {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	switch runtime.GOOS {
	case "linux", "darwin":
		return setupGitLinux()
	}
	return nil
	// pip3 config list

}
