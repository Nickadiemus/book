package visit

import (
	"fmt"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	fmt.Printf(" n.type: %v |\n n.Data: %v |\n n.Attr: %v |\n n.DataAtom: %v |\n n.Namespace: %v\n\n", n.Type, n.Data, n.Attr, n.DataAtom, n.Namespace)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}

// visit appends to links each link found in n and returns the result.
func VisitR(links []string, n *html.Node) []string {
	fmt.Printf(" n.type: %v |\n n.Data: %v |\n n.Attr: %v |\n n.DataAtom: %v |\n n.Namespace: %v\n\n", n.Type, n.Data, n.Attr, n.DataAtom, n.Namespace)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	fmt.Printf(" n.Parent: %v |\n n.FirstChild: %v |\n n.LastChild: %v |\n n.PrevSibling: %v |\n n.NextSibling: %v\n\n", n.Parent, n.FirstChild, n.LastChild, n.PrevSibling, n.NextSibling)
	fmt.Println("=========================================================================")

	if n.FirstChild != nil {
		links = VisitR(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = VisitR(links, n.NextSibling)
	}
	return links
}
