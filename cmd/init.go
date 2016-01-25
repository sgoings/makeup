// Copyright © 2015 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"os"
	"io/ioutil"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var include_lines = `# makeup-managed:begin
include makeup.mk
# makeup-managed:end

`

var makeupmk_contents = `# makeup-managed:begin
MAKEUP_DIR := .makeup

SUBMODULE_UPDATE := \$(shell git submodule update --init --recursive)
# makeup-managed:end
`

func safePrependWrite(filename string, ensure_contents string) {
	var orig string

	if _, err := os.Stat(filename); err == nil {
		orig_contents, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		orig = string(orig_contents)
	}

	if ! strings.Contains(orig, strings.TrimSpace(ensure_contents)) {
		output := fmt.Sprint(ensure_contents, orig)
		ioutil.WriteFile(filename, []byte(output), 0644)
	}
}

func WriteMakeupIncludeLines(makefile string) {
	safePrependWrite(makefile, include_lines)
	log.Printf("[INFO] injection logic: %s\n", makefile)
}

func WriteMakeupBootstrapFile(makeup_bootstrap_file string) {
	safePrependWrite(makeup_bootstrap_file, makeupmk_contents)
	log.Printf("[INFO] bootstrap file: %s\n", makeup_bootstrap_file)
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize makeup functionality in this repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		WriteMakeupBootstrapFile("makeup.mk")
		WriteMakeupIncludeLines("Makefile")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
