package setup

import (
	"fmt"
	"log"

	"autopard.com/desa/common"
)

func SetupNpmProxy(mirror string) error {
	cmds := []common.SysCmd{
		{Cmd: "npm", Params: []string{"config", "set", "registry", mirror}},
	}

	out, err := common.ExecuteCmds(cmds)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}
