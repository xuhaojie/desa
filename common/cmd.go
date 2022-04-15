package common

import (
	"log"
	"os/exec"
)

type SysCmd struct {
	Cmd    string
	Params []string
}

func (cmd *SysCmd) Execute() ([]byte, error) {
	command := exec.Command(cmd.Cmd, cmd.Params...)
	out, err := command.CombinedOutput()
	if err != nil {
		//fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("command.Run() failed with %s\n", err)
	}
	//fmt.Printf("combined out:\n%s\n", string(out))
	return out, err
}

func ExecuteCmds(cmds []SysCmd) ([]byte, error) {
	for _, cmd := range cmds {
		_, err := cmd.Execute()
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
