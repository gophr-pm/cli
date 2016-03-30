package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path/filepath"
)

const basicSkeleton string = "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main () {\n\tfmt.Println(\"hello world!\")\n}"

func RunInitCommand(goPath string, repoAuthor string, projectName string) {
	initPath := filepath.Join(goPath, "src", "github.com", repoAuthor, projectName)
	os.MkdirAll(initPath, 0777)

	// Now we need to glob to make sure a file name like that doesn't already exists
	fls, err := filepath.Glob(initPath + "/*.go")
	Check(err)

	if len(fls) > 0 {
		// check if the .go file names match your project name
		for _, fileName := range fls {
			if fileName == initPath+"/"+projectName+".go" {
				red := color.New(color.FgRed).SprintFunc()
				magenta := color.New(color.FgMagenta).SprintFunc()
				fmt.Printf("%s gophr %s %s file with that name already exists\n", red("✗"), red("ERROR"), magenta("init"))
				os.Exit(3)
			}
		}
	} else {
		// TODO throw error or gracefully exit
	}

	newFileBuffer := []byte(basicSkeleton)
	err = ioutil.WriteFile(filepath.Join(initPath, projectName)+".go", newFileBuffer, 0644)
	Check(err)
}
