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
	"fmt"
	"os/exec"
	"strings"
	"log"
	"path"
	"path/filepath"

	"github.com/deis/makeup/cmd/bag"
	"github.com/spf13/cobra"
)

func GetKitName(repo_url string) string {
	repo_basename := path.Base(repo_url)
	return strings.TrimSuffix(repo_basename, filepath.Ext(repo_basename))
}

func KitExists(repo_name string) bool {
	output, err := exec.Command("git", "submodule", "status").CombinedOutput()
	if err != nil {
		log.Fatalf("[ERROR] git submodule status failed with:\n%s\n", output)
	}

	if strings.Contains(string(output), repo_name) {
		return true
	}
	return false
}

func AddSubmodule(repo_path string) {
	repo_name := GetKitName(repo_path)

	if KitExists(repo_name) {
		log.Printf("[DEBUG] git submodule %s already exists!", repo_name)
	} else {
		url := fmt.Sprint("https://", repo_path, ".git")
		path := fmt.Sprint(bag.SubmoduleDir, "/", repo_name)
		output, err := exec.Command("git", "submodule", "add", url, path).CombinedOutput()
		if err != nil {
			log.Fatalf("[ERROR] git submodule add failed with:\n%s\n", output)
		}
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a makeup kit to this project",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			AddSubmodule(args[0])
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
