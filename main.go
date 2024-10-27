package main

import (
	"flag"
	"log"
	"structure/utils"
)

func main() {

	// Define the command-line flags
	hereFlag := flag.Bool("here", false, "Create project structure in the current directory")
	projectRootName := flag.String("project-root", "", "Name of the project root directory")
	structFile := flag.String("structure", "structure.yaml", "YAML file containing the project structure")

	flag.Parse()
	// Parse the YAML structure
	root, err := utils.ParseYAMLStructure(*structFile, *projectRootName)
	if err != nil {
		log.Fatalf("\033[1;33mError parsing structure: %v\033[0m", err)
	}

	var parentDir string
	if *hereFlag {
		// If --here is set, use the current directory
		parentDir = "."
		// Create children directly in the current directory
		for _, child := range root.Children {
			utils.MakeDirectoryAndFilesTree(child, parentDir)
		}
	} else {
		// If --here is not set, require a project root name
		if *projectRootName == "" {
			log.Fatalf("\033[1;33mError: project root name is required when --here is not specified\033[0m")
		}

		err = utils.MakeDirectoryAndFilesTree(root, parentDir)
		if err != nil {
			log.Fatalf("\033[1;33mError creating directory and files: %v\033[0m", err)
		}
	}

	if err != nil {
		log.Fatalf("\033[1;33mError creating directory and files: %v\033[0m", err)
	}
}
