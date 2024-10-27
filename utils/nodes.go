package utils

type Node struct {
	Name     string
	IsFile   bool
	Children []*Node
}

// CreateNode creates a new Node (file or directory)
func CreateNode(name string, isFile bool) *Node {
	return &Node{Name: name, IsFile: isFile, Children: []*Node{}}
}

// AddChild adds a child node to a directory node
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}
