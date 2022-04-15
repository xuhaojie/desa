package clean

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"autopard.com/desa/common"
)

func cleanGoProject(projectDir string) error {
	cmds := []common.SysCmd{
		{Cmd: "go", Params: []string{"clean"}},
	}
	os.Chdir(projectDir)
	out, err := common.ExecuteCmds(cmds)
	fmt.Println(string(out))
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return nil
}

func searchGoProjects(dir string) []string {
	var projtects []string

	err := filepath.Walk(dir, func(dir string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		} else {
			if f.Name() == "go.mod" {
				projectDir := filepath.Dir(dir)
				projtects = append(projtects, projectDir)
				fmt.Println("find project", projectDir)
			}
		}
		//println(path)
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return projtects
}

func CleanGoProjects(path string) error {
	projects := searchGoProjects(path)
	for _, p := range projects {
		fmt.Println("cleaning project", p, "...")
		cleanGoProject(p)
	}
	return nil
}
