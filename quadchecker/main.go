package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func makeQuad(x, y int, tl, tr, bl, br, hor, ver string) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var b strings.Builder

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0 && j == 0:
				b.WriteString(tl)
			case i == 0 && j == x-1:
				b.WriteString(tr)
			case i == y-1 && j == 0:
				b.WriteString(bl)
			case i == y-1 && j == x-1:
				b.WriteString(br)
			case i == 0 || i == y-1:
				b.WriteString(hor)
			case j == 0 || j == x-1:
				b.WriteString(ver)
			default:
				b.WriteString(" ")
			}
		}
		if i != y-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Not a quad function")
		return
	}

	input := strings.TrimSpace(string(data))
	if input == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(input, "\n")

	// validate rectangle shape
	width := len(lines[0])
	for _, l := range lines {
		if len(l) != width {
			fmt.Println("Not a quad function")
			return
		}
	}
	height := len(lines)

	type quad struct {
		name string
		val  string
	}

	quads := []quad{
		{"quadA", makeQuad(width, height, "o", "o", "o", "o", "-", "|")},
		{"quadB", makeQuad(width, height, "/", "\\", "\\", "/", "*", "*")},
		{"quadC", makeQuad(width, height, "A", "A", "C", "C", "B", "B")},
		{"quadD", makeQuad(width, height, "A", "C", "A", "C", "B", "B")},
		{"quadE", makeQuad(width, height, "A", "C", "C", "A", "B", "B")},
	}

	results := []string{}

	for _, q := range quads {
		if input == q.val {
			results = append(results,
				fmt.Sprintf("[%s] [%d] [%d]", q.name, width, height))
		}
	}

	if len(results) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	fmt.Println(strings.Join(results, " || "))
}
