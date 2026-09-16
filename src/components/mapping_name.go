package components

import (
	"fmt"
	"path"
	"smolmap/src/tiny"
	"strings"
)

var (
	TEXT_UNKNOWN    = format("unknown", COLOR_GRAY, STYLE_BOLD)
	TEXT_NOT_MAPPED = format("not mapped to anything", COLOR_YELLOW, STYLE_BOLD)
)

func RenderMappingName(sb *strings.Builder, mapping *tiny.Mapping, db *tiny.MappingDatabase) string {
	switch mapping.Type {
	case tiny.M_TYPE_ClASS:
		renderClassInfo(sb, mapping)
	case tiny.M_TYPE_FIELD:
		renderFieldInfo(sb, mapping, db)
	case tiny.M_TYPE_METHOD:
		renderFieldInfo(sb, mapping, db)
	}
	return sb.String()
}

func renderClassInfo(sb *strings.Builder, mapping *tiny.Mapping) {
	classNameFull := mapping.Names["named"]
	intermediaryClassName := mapping.Names["intermediary"]
	officialName := mapping.Names["official"]

	fmt.Fprintf(sb, "%s from %s\n",
		format(path.Base(classNameFull), COLOR_MAGENTA, STYLE_BOLD),
		format(classNameFull, COLOR_CYAN, STYLE_BOLD),
	)

	renderSubSection(sb, "Official (obfuscated) name", format(officialName, COLOR_CYAN, STYLE_BOLD))
	renderSubSection(sb, "Intermediary              ", format(intermediaryClassName, COLOR_CYAN, STYLE_BOLD))
}

func renderFieldInfo(sb *strings.Builder, mapping *tiny.Mapping, db *tiny.MappingDatabase) {
	intermediaryName := mapping.Names["intermediary"]
	methodName := mapping.Names["named"]
	officialName := mapping.Names["official"]

	fromClass := db.FindClassByFieldOrMethod(methodName)
	formattedClassPath := formatClassPath(fromClass)

	if intermediaryName == methodName {
		fmt.Fprintf(sb, "%s is %s inside %s\n",
			format(intermediaryName, COLOR_MAGENTA, STYLE_BOLD),
			TEXT_NOT_MAPPED,
			formattedClassPath,
		)
	} else {
		fmt.Fprintf(sb, "%s is %s from %s\n",
			format(intermediaryName, COLOR_MAGENTA, STYLE_BOLD),
			format(methodName, COLOR_BLUE, STYLE_BOLD),
			formattedClassPath,
		)

		renderSubSection(sb, "Descriptor                ", format(mapping.Descriptor, COLOR_GREEN, STYLE_BOLD))
	}
	renderSubSection(sb, "Official (obfuscated) name", format(officialName, COLOR_CYAN, STYLE_BOLD))
	renderSubSection(sb, "Intermediary              ", format(intermediaryName, COLOR_CYAN, STYLE_BOLD))

	if fromClass != nil {
		addNewLine(sb)
		renderClassInfo(sb, fromClass)
	}

	addNewLine(sb)

}

func renderMethodInfo(sb *strings.Builder, mapping *tiny.Mapping, db *tiny.MappingDatabase) {
	intermediaryName := mapping.Names["intermediary"]
	methodName := mapping.Names["named"]
	officialName := mapping.Names["official"]

	fromClass := db.FindClassByFieldOrMethod(methodName)
	formattedClassPath := formatClassPath(fromClass)

	if intermediaryName == methodName {
		fmt.Fprintf(sb, "%s is %s inside %s\n",
			format(intermediaryName, COLOR_CYAN, STYLE_BOLD),
			TEXT_NOT_MAPPED,
			formattedClassPath,
		)
	} else {
		fmt.Fprintf(sb, "%s is %s` from %s\n",
			format(intermediaryName, COLOR_MAGENTA, STYLE_BOLD),
			format(methodName, COLOR_YELLOW, STYLE_BOLD),
			formattedClassPath,
		)

		renderSubSection(sb, "Descriptor                ", format(mapping.Descriptor, COLOR_GREEN, STYLE_BOLD))
	}

	if fromClass != nil {
		addNewLine(sb)
		renderClassInfo(sb, fromClass)
	}

	renderSubSection(sb, "Official (obfuscated) name", format(officialName, COLOR_CYAN, STYLE_BOLD))
}

func renderSubSection(sb *strings.Builder, name string, stuff string) {
	fmt.Fprintf(sb, "╰ %s : %s\n", format(name, COLOR_GRAY), stuff)
}

func addNewLine(sb *strings.Builder) {
	fmt.Fprintln(sb)
}

func formatClassPath(classInfo *tiny.Mapping) string {
	if classInfo == nil {
		return TEXT_UNKNOWN
	}
	return format(classInfo.Names["named"], COLOR_CYAN, STYLE_BOLD)
}
