package builder

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// AtomXMLBuilder represents the object that is able to construct the XML content in ATOM feed standard
type AtomXMLBuilder struct {
	rootNode *XMLNode // root of the whole XML
	feedNode *XMLNode // working node for feed
}

// String is the resulting assembly of the builder (implements Stringer interface)
func (builder *AtomXMLBuilder) String() string {
	return builder.rootNode.String()
}

// AddChild is a general interface for adding the child nodes to the feed
func (builder *AtomXMLBuilder) AddChild(child *XMLNode) {
	builder.feedNode.AddChild(child)
}

// AddFeedSubTitle is a specific builder's method to add the subtitle node to the feed
func (builder *AtomXMLBuilder) AddFeedSubTitle(subtitle string) {
	builder.AddChild(&XMLNode{
		name: "subtitle",
		text: subtitle,
	})
}

// AddFeedAuthor is a specific builder's method to add the author of the feed, with name and email as parameters
func (builder *AtomXMLBuilder) AddFeedAuthor(name string, email string) {
	builder.AddChild(&XMLNode{
		name: "author",
		children: []*XMLNode{
			{
				name: "name",
				text: name,
			},
			{
				name: "email",
				text: email,
			},
		},
	})
}

// AddEntry is a specific builder's method to add entry to the feed with optional children and required title. Adds
// entry with the random UUID.
func (builder *AtomXMLBuilder) AddEntry(title string, children []*XMLNode) {
	uuid := uuid.NewV1()

	entryTitleNode := XMLNode{
		name: "title",
		text: title,
	}

	entryIdNode := XMLNode{
		name: "id",
		text: fmt.Sprintf("urn:uuid:%v", uuid),
	}

	entryUpdatedNode := XMLNode{
		name: "updated",
		text: time.Now().Format(time.RFC3339),
	}

	builder.AddChild(&XMLNode{
		name:     "entry",
		children: append([]*XMLNode{&entryTitleNode, &entryIdNode, &entryUpdatedNode}, children...),
	})
}

// NewAtomXMLBuilder is the builder initial constructor. Produces builder with the valid skeleton markup of required
// tags with random ID and defined title. This should be used for a public interface entrypoint.
func NewAtomXMLBuilder(title string) *AtomXMLBuilder {
	uuid := uuid.NewV1()

	idNode := XMLNode{
		name: "id",
		text: fmt.Sprintf("urn:uuid:%v", uuid),
	}

	updatedNode := XMLNode{
		name: "updated",
		text: time.Now().Format(time.RFC3339),
	}

	titleNode := XMLNode{
		name: "title",
		text: title,
	}

	versionNode := XMLNode{
		name:      "?xml",
		contained: true,
		attributes: map[string]string{
			"version":  "1.0",
			"encoding": "utf-8",
		},
	}

	feedNode := XMLNode{
		name: "feed",
		attributes: map[string]string{
			"xmlns": "http://www.w3.org/2005/Atom",
		},
		children: []*XMLNode{&titleNode, &idNode, &updatedNode},
	}

	return &AtomXMLBuilder{
		rootNode: &XMLNode{
			children: []*XMLNode{&versionNode, &feedNode},
		},
		feedNode: &feedNode,
	}
}
