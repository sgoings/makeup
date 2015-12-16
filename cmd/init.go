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

	"github.com/spf13/cobra"
)

func WriteMakeupIncludeLines() {
	file, err := os.Create("Makefile")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	include_script := `MAKEUP_DIR := .makeup

SUBMODULE_UPDATE := \$(shell git submodule update --init --recursive)

`

	num_bytes, err := file.WriteString(include_script)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %d bytes to makeup.mk script", num_bytes)
}

func WriteMakeupBootstrapFile() {
	file, err := os.Create("makeup.mk")
	if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

	include_script := `MAKEUP_DIR := .makeup

SUBMODULE_UPDATE := \$(shell git submodule update --init --recursive)

`

	num_bytes, err := file.WriteString(include_script)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %d bytes to makeup.mk script", num_bytes)
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize makeup, creating a makeup.yaml file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		WriteMakeupBootstrapFile()
		WriteMakeupIncludeLines()
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
