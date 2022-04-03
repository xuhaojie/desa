package env

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
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

func setupGolangLinux() error {
	//$ go env -w GO111MODULE=on
	//$ go env -w GOPROXY=https://goproxy.cn,directory
	params1 := []string{"env", "-w", "GO111MODULE=on"}
	out, err := execCommand("go", params1)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}

	params2 := []string{"env", "-w", "GOPROXY=https://goproxy.cn,directory"}
	out, err = execCommand("go", params2)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}
	return nil
}

func SetupGolangProxy() error {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	switch runtime.GOOS {
	case "linux":
		return setupGolangLinux()
	}
	return nil
	// pip3 config list

}
