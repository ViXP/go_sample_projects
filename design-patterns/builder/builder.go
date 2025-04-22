// Package builder implements the Builder design pattern.
// The struct AtomXMLBuilder represents the object that is able to construct the XML content in ATOM feed standard
// By providing the API for the manual decision on the nodes structure and their contents.
package builder

import (
	"fmt"
)

func Run() {
	builder := NewAtomXMLBuilder("The ATOM feed title")
	builder.AddFeedSubTitle("The ATOM feed subtitle")
	builder.AddFeedAuthor("The Author", "author@mail.com")
	builder.AddEntry("First entry in the ATOM feed", []*XMLNode{
		{
			name: "summary",
			text: "The optional summary of the first entry",
		},
	}).AddEntry("Second simple entry in the ATOM feed", []*XMLNode{}) // add other in the chain manner

	fmt.Println(builder)
}
