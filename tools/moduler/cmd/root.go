// Copyright 2018 Microsoft Corporation
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
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/tools/apidiff/repo"
	"github.com/spf13/cobra"
)

const (
	tagPrefix         = "// tag: "
	versionFile       = "version.go"
	profileFolder     = "profiles"
	messageForProfile = "Update profiles"
)

var rootCmd = &cobra.Command{
	Use:   "module-releaser <sdk root path>",
	Short: "Release a new module by pushing new tags to github",
	Long: `This tool search the whole SDK folder for tags in version.go files which are produced by the versioner tool, 
and push the new tags to github.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return theCommand(args)
	},
}

var (
	getTagsHook func(string) ([]string, error)
	dryRunFlag  bool
	verboseFlag bool
)

func init() {
	getTagsHook = getTags
	rootCmd.PersistentFlags().BoolVarP(&dryRunFlag, "dry-run", "d", false, "dry run, only list the detected tags, do not add or push them")
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "verbose output")
}

// Execute the tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func theCommand(args []string) error {
	path, err := filepath.Abs(args[0])
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}
	// update local repo
	if err := updateLocalRepo(path); err != nil {
		return fmt.Errorf("failed to update local repo: %v", err)
	}
	// get list of all existing tags
	tags, err := getTagsHook(path)
	if err != nil {
		return fmt.Errorf("failed to list tags: %v", err)
	}
	vprintf("Get %d tags from remote\n", len(tags))
	// find all version.go files
	files, err := findAllFiles(path, versionFile)
	if err != nil {
		return fmt.Errorf("failed to find all version.go files: %v", err)
	}
	// read new tags and push them in version.go
	if err := readAndAddTags(tags, files); err != nil {
		return fmt.Errorf("failed during reading and pushing new tags: %v", err)
	}
	if !dryRunFlag {
		// push new tags
		if err := pushNewTags(); err != nil {
			return fmt.Errorf("failed to push tags: %v", err)
		}
		// generate profiles
		profilePath := filepath.Join(path, profileFolder)
		if err := generateProfiles(profilePath); err != nil {
			return fmt.Errorf("failed during generating profiles: %v", err)
		}
		// push repo
		if err := pushRepo(profilePath, messageForProfile); err != nil {
			return fmt.Errorf("failed during add and commit new files: %v", err)
		}
	}
	return nil
}

func updateLocalRepo(repoPath string) error {
	return executeCommand("git", "pull", "--tags")
}

func getTags(repoPath string) ([]string, error) {
	wt, err := repo.Get(repoPath)
	if err != nil {
		return nil, err
	}
	return wt.ListTags("*")
}

func tagExists(tags []string, tag string) bool {
	for _, item := range tags {
		if item == tag {
			return true
		}
	}
	return false
}

func findAllFiles(root, filename string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && info.Name() == filename {
			files = append(files, path)
			return nil
		}
		return nil
	})
	return files, err
}

func readAndAddTags(tags []string, files []string) error {
	for _, file := range files {
		newTag, err := findNewTagInFile(file)
		if err != nil {
			return fmt.Errorf("failed to read tag in file '%s': %v", file, err)
		}
		if newTag == "" {
			printf("Found empty tag in file '%s'\n", file)
			continue
		} else {
			vprintf("Found tag '%s' in file '%s'\n", newTag, file)
		}
		if tagExists(tags, newTag) {
			vprintf("Tag '%s' already exists, skipping", newTag)
			continue
		}
		// add new tag to list
		tags = append(tags, newTag)
		// add new tag to git
		if err := addNewTag(newTag); err != nil {
			return fmt.Errorf("failed to push tag '%s': %v", newTag, err)
		}
	}
	return nil
}

func findNewTagInFile(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	tag := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, tagPrefix) {
			tag = strings.ReplaceAll(line, tagPrefix, "")
			break
		}
	}
	return tag, nil
}

func addNewTag(newTag string) error {
	if dryRunFlag {
		vprintf("Found new tag '%s'\n", newTag)
	} else {
		vprintf("Adding new tag '%s'\n", newTag)
		if err := executeCommand("git", "tag", "-a", newTag, "-m", newTag); err != nil {
			return err
		}
	}
	return nil
}

func pushNewTags() error {
	vprintln("Pushing new tags")
	return executeCommand("git", "push", "origin", "--tags")
}

func generateProfiles(profilePath string) error {
	vprintf("Running `go generate` in %s\n", profilePath)
	return executeCommand("go", "generate", profilePath)
}

func pushRepo(profilePath, message string) error {
	vprintf("Add files in %s\n", profilePath)
	if err := executeCommand("git", "add", profilePath); err != nil {
		return err
	}
	if err := executeCommand("git", "commit", "-m", message); err != nil {
		return err
	}
	if err := executeCommand("git", "push"); err != nil {
		return err
	}
	return nil
}

func executeCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.New(string(output))
	}
	return nil
}
