package count

import "golang.org/x/net/html"

func CountElements(count int, element string, n *html.Node) int {
	if n.Type == html.ElementNode && element == n.Data {
		count++
	}
	if n.FirstChild != nil {
		count = CountElements(count, element, n.FirstChild)
	}
	if n.NextSibling != nil {
		count = CountElements(count, element, n.NextSibling)
	}
	return count
}
