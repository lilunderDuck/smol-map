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
	flag.StringVar(&tinyMappingFile, "tiny", "<no .tiny file>", "path to .tiny v1 mapping")
	flag.Parse()

	root := views.NewSearchNameModel(tinyMappingFile)

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
