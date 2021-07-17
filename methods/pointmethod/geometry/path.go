package geometry

import "fmt"

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	fmt.Println("path: %v", path)
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
			fmt.Printf("index: %v\tpath[i-1]: %v\tpath[i]: %v\tsum: %v\n", i, path[i-1], path[i], sum)

		}
	}
	return sum
}
