package builder

import (
	"fmt"
	"strings"
)

// XMLNode is a general "low level" struct used to creating the nodes tree
type XMLNode struct {
	name       string
	text       string
	children   []*XMLNode
	attributes map[string]string
	contained  bool
}

// String is a stringified representation of the XMLNode including it's name, children and attributes
func (node *XMLNode) String() (result string) {
	var childrenBuilder strings.Builder
	var attributesBuilder strings.Builder

	for _, child := range node.children {
		childrenBuilder.WriteString(child.String())
	}

	for attrib, value := range node.attributes {
		attributesBuilder.WriteString(fmt.Sprintf(" %s=\"%s\"", attrib, value))
	}

	if node.text != "" {
		childrenBuilder.Reset()
		childrenBuilder.WriteString(node.text)
	}

	if node.name != "" {
		if !node.contained {
			result = fmt.Sprintf("<%s%s>%s</%s>", node.name, attributesBuilder.String(), childrenBuilder.String(), node.name)
		} else {
			result = fmt.Sprintf("<%s%s/>", node.name, attributesBuilder.String())
		}
	} else {
		result = childrenBuilder.String()
	}
	return
}

// AddChild is a general interface for adding the children of the XMLNode
func (node *XMLNode) AddChild(child *XMLNode) *XMLNode {
	node.children = append(node.children, child)
	return node
}
