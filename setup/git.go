package setup

import (
	"fmt"
	"log"

	"autopard.com/desa/common"
)

func SetupGit() error {
	cmds := []common.SysCmd{
		{Cmd: "git", Params: []string{"config", "--global", "user.name", "xuhaojie"}},
		{Cmd: "git", Params: []string{"config", "--global", "user.email", "xuhaojie@hotmail.com"}},
	}

	out, err := common.ExecuteCmds(cmds)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}
