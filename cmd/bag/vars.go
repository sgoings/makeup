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
	log.Printf("listing vars for bag %s", path)
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
		if variable.description != "" {
			fmt.Println(variable.name, " - ", variable.description)
		} else {
			fmt.Println(variable.name)
		}
	}
}

func AddVarsCommand(cmd *cobra.Command) {
	cmd.AddCommand(varsCmd)
}
