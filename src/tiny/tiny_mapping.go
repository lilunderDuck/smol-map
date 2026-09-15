package tiny

import (
	"bufio"
	"fmt"
	"os"
	"path"
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

func DetectForTinyV1(file string, scannerOut *bufio.Scanner) error {
	r, err := os.Open(file)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return fmt.Errorf("empty mapping file")
	}

	header := strings.Split(scanner.Text(), "\t")
	if len(header) == 1 { // Handle space-separated v1 header
		header = strings.Fields(scanner.Text())
	}

	if scannerOut != nil {
		scannerOut = scanner
	}

	// Detect version from header
	if header[0] != "v1" {
		return fmt.Errorf("unsupported tiny format header: %s", header[0])
	}

	return nil
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

	if header[0] != "v1" {
		return nil, fmt.Errorf("unsupported tiny format header: %s", header[0])
	}

	m := &Mapping{}
	return parseTinyV1(scanner, m)
}

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

func (m *Mapping) FindClass(query string) (*ClassMapping, bool) {
	for i := range m.Classes {
		class := &m.Classes[i]

		for _, name := range class.Names {
			// Direct full match (e.g., "net/minecraft/class_155")
			if name == query {
				return class, true
			}

			// Base name match (e.g., matching "class_155" inside "net/minecraft/class_155")
			if path.Base(name) == query {
				return class, true
			}
		}
	}
	return nil, false
}

func (c *ClassMapping) GetNamespacesMap(namespaces []string) map[string]string {
	result := make(map[string]string)
	for i, ns := range namespaces {
		if i < len(c.Names) {
			result[ns] = c.Names[i]
		}
	}
	return result
}
