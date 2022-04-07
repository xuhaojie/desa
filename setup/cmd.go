package setup

import (
	"log"
	"os/exec"
)

type SysCmd struct {
	cmd    string
	params []string
}

func (cmd *SysCmd) execute() ([]byte, error) {
	command := exec.Command(cmd.cmd, cmd.params...)
	out, err := command.CombinedOutput()
	if err != nil {
		//fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("command.Run() failed with %s\n", err)
	}
	//fmt.Printf("combined out:\n%s\n", string(out))
	return out, err
}

func executeCmds(cmds []SysCmd) ([]byte, error) {
	for _, cmd := range cmds {
		_, err := cmd.execute()
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
