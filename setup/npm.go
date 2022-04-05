package setup

import (
	"fmt"
	"log"
)

func SetupNpmProxy(mirror string) error {
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
