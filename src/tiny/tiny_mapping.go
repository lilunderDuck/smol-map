package tiny

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mapping struct {
	Namespaces []string
	Classes    []ClassMapping
}

type ClassMapping struct {
	Names   []string
	Fields  []MemberMapping
	Methods []MemberMapping
}

type MemberMapping struct {
	Descriptor string
	Names      []string
}

func ParseTiny(file string) (*Mapping, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(r)

	if !scanner.Scan() {
		return nil, fmt.Errorf("empty mapping file")
	}

	header := strings.Split(scanner.Text(), "\t")
	if len(header) == 1 { // Handle space-separated v1 header
		header = strings.Fields(scanner.Text())
	}

	m := &Mapping{}

	// Detect version from header
	switch header[0] {
	case "v1":
		m.Namespaces = header[1:]
		return parseTinyV1(scanner, m)
	default:
		return nil, fmt.Errorf("unsupported tiny format header: %s", header[0])
	}
}

// Parse flat Tiny V1 format
func parseTinyV1(scanner *bufio.Scanner, m *Mapping) (*Mapping, error) {
	var currentClass *ClassMapping

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, "\t")
		switch parts[0] {
		case "CLASS":
			m.Classes = append(m.Classes, ClassMapping{
				Names: parts[1:],
			})
			currentClass = &m.Classes[len(m.Classes)-1]

		case "FIELD":
			if currentClass != nil {
				currentClass.Fields = append(currentClass.Fields, MemberMapping{
					Descriptor: parts[1],
					Names:      parts[2:],
				})
			}

		case "METHOD":
			if currentClass != nil {
				currentClass.Methods = append(currentClass.Methods, MemberMapping{
					Descriptor: parts[1],
					Names:      parts[2:],
				})
			}
		}
	}
	return m, scanner.Err()
}
