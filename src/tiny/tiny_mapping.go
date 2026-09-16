package tiny

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	M_TYPE_ClASS  = "CLASS"
	M_TYPE_FIELD  = "FIELD"
	M_TYPE_METHOD = "METHOD"
)

type Mapping struct {
	Type       string            `json:"type"`
	Location   string            `json:"location,omitempty"`
	Descriptor string            `json:"descriptor,omitempty"`
	Names      map[string]string `json:"names"`
}

type MappingDatabase struct {
	Namespaces []string
	Lookup     map[string][]Mapping
}

func NewMappingDatabase() *MappingDatabase {
	return &MappingDatabase{
		Lookup: make(map[string][]Mapping),
	}
}

func (db *MappingDatabase) Index(m Mapping) {
	for _, name := range m.Names {
		if name == "" {
			continue
		}

		// Extract short name from path if applicable (e.g. "class_155" from "net/minecraft/class_155")
		shortName := name
		if idx := strings.LastIndex(name, "/"); idx != -1 {
			shortName = name[idx+1:]
		}

		// add full identifier path to lookup
		db.Lookup[name] = append(db.Lookup[name], m)

		// add short name key if different (e.g. "class_155 -> [...]")
		if shortName != name {
			db.Lookup[shortName] = append(db.Lookup[shortName], m)
		}
	}
}

func ParseTiny(fileLocation string) (*MappingDatabase, error) {
	r, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	db := NewMappingDatabase()
	scanner := bufio.NewScanner(r)

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Skip comments or empty lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, "\t")

		// Parse Header Line: v1 <namespace1> <namespace2> ...
		if lineNumber == 1 || parts[0] == "v1" {
			if parts[0] != "v1" {
				return nil, fmt.Errorf("invalid header: expected v1 prefix")
			}
			db.Namespaces = parts[1:]
			continue
		}

		kind := parts[0]
		namesMap := make(map[string]string)

		switch kind {
		case M_TYPE_ClASS:
			// CLASS <ns1_name> <ns2_name> ...
			for i, ns := range db.Namespaces {
				if i+1 < len(parts) {
					namesMap[ns] = parts[i+1]
				}
			}

			m := Mapping{
				Type:  M_TYPE_ClASS,
				Names: namesMap,
			}
			db.Index(m)

		case M_TYPE_FIELD:
			// FIELD <enclosing_class> <descriptor> <ns1_name> <ns2_name> ...
			if len(parts) < 3+len(db.Namespaces) {
				continue
			}
			owner := parts[1]
			desc := parts[2]
			for i, ns := range db.Namespaces {
				namesMap[ns] = parts[3+i]
			}

			m := Mapping{
				Type:       M_TYPE_FIELD,
				Location:   owner,
				Descriptor: desc,
				Names:      namesMap,
			}
			db.Index(m)

		case M_TYPE_METHOD:
			// METHOD <enclosing_class> <descriptor> <ns1_name> <ns2_name> ...
			if len(parts) < 3+len(db.Namespaces) {
				continue
			}
			owner := parts[1]
			desc := parts[2]
			for i, ns := range db.Namespaces {
				namesMap[ns] = parts[3+i]
			}

			m := Mapping{
				Type:       M_TYPE_METHOD,
				Location:   owner,
				Descriptor: desc,
				Names:      namesMap,
			}
			db.Index(m)
		}
	}

	return db, scanner.Err()
}

func (db *MappingDatabase) FindClassByFieldOrMethod(methodOrField string) *Mapping {
	for _, methodMapping := range db.Lookup[methodOrField] {
		parentClassName := methodMapping.Location
		if parentClassMappings, found := db.Lookup[parentClassName]; found {
			for _, classMapping := range parentClassMappings {
				if classMapping.Type == "CLASS" {
					return &classMapping
				}
			}
		}
	}

	return nil
}
