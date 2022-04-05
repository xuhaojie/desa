package setup

import (
	"fmt"
	"log"
)

func SetupGolangProxy(mirror string) error {
	//$ go env -w GO111MODULE=on
	//$ go env -w GOPROXY=https://goproxy.cn,direct
	proxy := fmt.Sprintf("GOPROXY=%s,direct", mirror)
	cmds := []Cmd{
		{cmd: "go", params: []string{"env", "-w", "GO111MODULE=on"}},
		{cmd: "go", params: []string{"env", "-w", proxy}},
	}
	out, err := executeCmds(cmds)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}
