package main

import (
	"flag"
	"fmt"
	"os"
	"smolmap/src/views"

	tea "charm.land/bubbletea/v2"
)

func main() {
	var tinyMappingFile string
	flag.StringVar(&tinyMappingFile, "tiny", "", "path to .tiny v1 mapping")
	flag.Parse()

	root := views.NewRootModel()
	if tinyMappingFile == "" {
		root.GoTo(views.VIEW_SEARCH_NAME)
	}

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Enter any intermediary name:")
	// if scanner.Scan() {
	// 	input := scanner.Text()
	// 	fmt.Printf("input value is: %s!\n", input)
	// }
	// fmt.Println(components.MappingName("class_155", "SharedConstants", "net/minecraft/SharedConstants"))
	p := tea.NewProgram(root)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
