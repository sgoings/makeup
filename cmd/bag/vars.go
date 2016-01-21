package bag

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	dependencies  = "DEPENDENCIES"
	varListPrefix = "#		- "
)

func lineCommented(text string) string {
	return "# " + text
}

var varsCmd = &cobra.Command{
	Use:   "vars",
	Short: "List the input variables for a bag or individual Makefile inside a bag",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Printf("[ERROR] Bag or individual makefile name not given")
			os.Exit(1)
		}
		if strings.LastIndex(args[0], ".mk") != -1 {
			listMakefileVars(args[0])
		} else {
			listBagVars(args[0])
		}
	},
}

func listBagVars(path string) {
	makefiles := make(map[string][]variable)
	makefileErrs := make(map[string]error)
	spl := strings.Split(path, "/")
	if len(spl) != 3 {
		log.Printf("[ERROR] bag path %s is invalid", path)
		os.Exit(1)
	}
	bagName := spl[len(spl)-1]
	relPath := filepath.Join(SubmoduleDir, bagName)
	err := filepath.Walk(relPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && filepath.Base(path) == ".git" {
			return filepath.SkipDir
		} else if filepath.Base(path) == ".git" || info.IsDir() {
			return nil
		}

		fd, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fd.Close()
		vars, err := parseVars(fd)
		if err != nil {
			makefileErrs[path] = err
		} else {
			makefiles[path] = vars
		}
		return nil
	})
	if err != nil && err != filepath.SkipDir {
		log.Printf("[ERROR] walking the bag directory (%s)", err)
		os.Exit(1)
	}

	for makefileName, vars := range makefiles {
		fmt.Println("-- ", filepath.Base(makefileName), " --")
		for _, v := range vars {
			fmt.Println(v.String())
		}
	}
}

func listMakefileVars(path string) {
	spl := strings.Split(path, "/")
	if len(spl) != 4 {
		log.Printf("[ERROR] makefile path %s is invalid", path)
		os.Exit(1)
	}
	bagName := spl[len(spl)-2]
	makefileName := spl[len(spl)-1]
	relPath := filepath.Join(SubmoduleDir, bagName, makefileName)
	fd, err := os.Open(relPath)
	if err != nil {
		log.Printf("[ERROR] opening makefile %s (%s)", path, err)
		os.Exit(1)
	}
	defer fd.Close()
	vars, err := parseVars(fd)
	if err != nil {
		log.Printf("[ERROR] parsing variables from %s (%s)", path, err)
		os.Exit(1)
	}
	for _, variable := range vars {
		fmt.Println(variable.String())
	}
}

func AddVarsCommand(cmd *cobra.Command) {
	cmd.AddCommand(varsCmd)
}
