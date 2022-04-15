package setup

import (
	"fmt"
	"log"

	"autopard.com/desa/common"
)

func SetupGolangProxy(mirror string) error {
	//$ go env -w GO111MODULE=on
	//$ go env -w GOPROXY=https://goproxy.cn,direct
	proxy := fmt.Sprintf("GOPROXY=%s,direct", mirror)
	cmds := []common.SysCmd{
		{Cmd: "go", Params: []string{"env", "-w", "GO111MODULE=on"}},
		{Cmd: "go", Params: []string{"env", "-w", proxy}},
	}
	out, err := common.ExecuteCmds(cmds)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}
