package setup

import (
	"fmt"
	"log"
)

func SetupGit() error {
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
