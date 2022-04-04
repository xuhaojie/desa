package common

import (
	"fmt"
	"log"
	"os/exec"
)

func execCommand(cmd string, params []string) ([]byte, error) {
	//$ go env -w GO111MODULE=on
	//$ go env -w GOPROXY=https://goproxy.cn,directory
	command := exec.Command(cmd, params...)
	out, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("command.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	return out, err
}
