package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterszarvas94/lytepage/pkg/utils"
	"github.com/peterszarvas94/lytepage/pkg/version"
)

func main() {
	version := version.Version

	fmt.Printf("Publishing version: %s\n", version)

	// checking tag
	ok, err := utils.RemoteTagExists(version)
	utils.CheckError(err, "Can not get remote tags")
	if !ok {
		fmt.Printf("Tag does not exists yet: %s \n", version)
	} else {
		fmt.Printf("Tag already exists: %s \n", version)
		os.Exit(1)
	}

	// check git status
	has, err := utils.HasUncomittedChanges()
	utils.CheckError(err, "Error checking uncommitted changes")
	if has {
		fmt.Println("You have uncomitted changes")
		os.Exit(1)
	} else {
		fmt.Println("You dont have uncomitted changes")
	}

	// changing go.mod files for templates:
	// - replacing lytepage version
	// - removing "replace" directives

	modFilePath := filepath.Join("scaffhold", "go.mod")

	modFile, err := os.Open(modFilePath)
	utils.CheckError(err, "No modfile found")

	var newContent strings.Builder
	scanner := bufio.NewScanner(modFile)

	for scanner.Scan() {
		line := scanner.Text()
		start := strings.Index(line, "github.com/peterszarvas94/lytepage")
		if start != -1 {
			parts := strings.Fields(line[start:])
			if len(parts) == 2 {
				line = line[:start] + parts[0] + " " + version
				fmt.Printf("Updated lytepage version number in file: %s\n", modFilePath)
			}
		}

		if strings.Contains(line, "replace github.com/peterszarvas94/lytepage") && !strings.HasPrefix(line, "// ") {
			fmt.Printf("Replace   : %s\n", line)
			newContent.WriteString("// ")
			newContent.WriteString(line)
			fmt.Printf("Commended out replace directive in file: %s\n", modFilePath)
		} else {
			fmt.Printf("Line is ok: %s\n", line)
			newContent.WriteString(line)
			newContent.WriteString("\n")
		}
	}

	err = scanner.Err()
	utils.CheckError(err, "Error reading modfile")

	err = os.WriteFile(modFilePath, []byte(newContent.String()), 0644)
	utils.CheckError(err, "Error writing modfile")

	err = modFile.Close()
	utils.CheckError(err, "Error closing modfile")

	fmt.Println("go.mod file has been updated and synchronized.")

	// git stuff
	err = utils.Cmd("git", "add", ".")
	utils.CheckError(err, "Error staging files")

	fmt.Println("Staged files")

	err = utils.Cmd("git", "commit", "-m", fmt.Sprintf("publish: %s", version))
	utils.CheckError(err, "Error committing files")

	fmt.Println("Commited files")

	err = utils.Cmd("git", "push")
	utils.CheckError(err, "Error pushing")

	fmt.Println("Pushed files")

	err = utils.Cmd("git", "tag", version, "-m", version)
	utils.CheckError(err, "Error tagging")

	err = utils.Cmd("git", "push", "--tags")
	utils.CheckError(err, "Error pushing tags")

	fmt.Println("Pushed tags")

	modFile, err = os.Open(modFilePath)
	utils.CheckError(err, "Error reopening modfile")
	defer modFile.Close()

	newContent.Reset()
	scanner = bufio.NewScanner(modFile)

	for scanner.Scan() {
		line := scanner.Text()
		start := strings.Index(line, "// replace github.com/peterszarvas94/lytepage")
		if start != -1 {
			line = strings.TrimPrefix(line, "// ")
			fmt.Printf("Replace directive restored in modfile: %s\n", modFilePath)
		}

		newContent.WriteString(line)
		newContent.WriteString("\n")
	}

	err = scanner.Err()
	utils.CheckError(err, "Error opening modfile")

	err = os.WriteFile(modFilePath, []byte(newContent.String()), 0644)
	utils.CheckError(err, "Error writing modfile")

	err = modFile.Close()
	utils.CheckError(err, "Error closing modfile")

	err = utils.Cmd("go", "mod", "tidy")
	utils.CheckError(err, "Error tidying")

	fmt.Printf("Tidied")
}
