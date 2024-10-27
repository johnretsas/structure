package utils

import (
	"bufio"
	"os"
	"strings"
)

// ParseYAMLStructure parses the YAML-like input to create a directory tree structure
func ParseYAMLStructure(filename string, projectRootName string) (*Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	root := CreateNode(projectRootName, false) // Create the root node
	stack := []*Node{root}                     // Stack to keep track of depth

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Calculate depth by counting leading spaces
		indentLevel := len(line) - len(strings.TrimLeft(line, " "))
		depth := indentLevel / 2

		// Determine if it's a file or directory and get the name
		line = strings.TrimSpace(line)
		isFile := !strings.HasSuffix(line, ":")
		name := strings.TrimSuffix(line, ":")

		if isFile {
			name = strings.TrimSuffix(name, ": file") // Remove " file" suffix if it's a file
		}

		// Create a new node
		node := CreateNode(name, isFile)

		// Adjust the stack to match the current depth
		if depth >= len(stack) {
			stack = append(stack, node)
		} else {
			stack = stack[:depth+1]
			stack[depth] = node
		}

		// Add the new node as a child to the previous node in the stack
		if depth > 0 {
			stack[depth-1].AddChild(node)
		} else {
			root.AddChild(node)
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return root, nil
}
