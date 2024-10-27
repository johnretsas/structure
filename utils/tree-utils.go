package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// PrintDirectoryTree recursively prints the directory tree structure
func PrintDirectoryTree(node *Node, depth int) {
	indent := strings.Repeat("  ", depth)

	if node.IsFile {
		log.Printf("%s- %s", indent, node.Name)
	} else {
		log.Printf("%s- %s/", indent, node.Name)
		for _, child := range node.Children {
			PrintDirectoryTree(child, depth+1)
		}
	}
}

func MakeDirectoryAndFilesTree(node *Node, parentDir string) error {
	// if node is a file create a file with the name of the node
	nodePath := filepath.Join(parentDir, node.Name)
	if node.IsFile {
		// create a file with the name of the node
		// and write some content to it
		file, err := os.Create(nodePath)
		if err != nil {
			log.Fatalf("Error creating file: %v", err)
			return err
		}
		defer file.Close()
	} else {
		// make a directory with the name of the node
		// 0755 is the permission for the directory
		// 0755 means the owner can read, write, and executes
		err := os.Mkdir(nodePath, 0755)
		if err != nil {
			log.Fatalf("Error creating directory: %v", err)
			return err
		}
		for _, child := range node.Children {
			MakeDirectoryAndFilesTree(child, nodePath)
		}
	}

	return nil
}
