package main

import (
	"flag"
	"fmt"
	"os"
	"piscine"
	"strings"
)

const (
	errorMess = "Usage: go run . [OPTION: COLOR] [OPTION: SUBSTRING] [STRING] \n or \nUsage: go run . [OPTION: COLOR] [OPTION: SUBSTRING] [STRING] [BANNER]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"\n\nFont options: Cross, Standard, Standard2, Shadow, Thinkertoy\nColor options: Red, Green, Yellow, Blue, Purple, Cyan, White, Pink, Orange, Grey, Brown, Beige."
)

func main() {
	text := ""
	substring := ""
	fontType := ""
	if len(os.Args) <= 1 {
		fmt.Println(errorMess)
		return
	}
	if len(os.Args) <= 1 {
		text = os.Args[1]
	} else if strings.HasPrefix(os.Args[1], "--") {
		if !strings.HasPrefix(os.Args[1], "--color=") {
			fmt.Println(errorMess)
			return
		}
	}
	color := flag.String("color", "reset", "Color of the substring")
	flag.Parse()
	// Define command-line flags

	if !piscine.Exist(color) {
		if len(os.Args) == 2 {
			*color = "reset"
		} else {
			fmt.Println("This color is not available.")
			fmt.Println(errorMess)
			return
		}
	}

	if len(flag.Args()) < 1 || len(flag.Args()) > 3 {
		fmt.Println(errorMess)
		return
	} else if len(flag.Args()) == 1 {
		text = flag.Args()[0]
		fontType = "standard"
	} else if len(flag.Args()) == 2 {
		if !piscine.FontType(strings.ToLower(flag.Args()[1])) {
			substring = flag.Args()[0]
			text = flag.Args()[1]
			fontType = "standard"
		} else {
			text = flag.Args()[0]
			fontType = flag.Args()[1]
		}
	} else if len(flag.Args()) == 3 {
		substring = flag.Args()[0]
		text = flag.Args()[1]
		fontType = flag.Args()[2]
	}

	if text == "" {
		fmt.Println(errorMess)
		return
	}

	if !piscine.FontType(strings.ToLower(fontType)) {
		fmt.Println(errorMess)
		return
	}

	filename := piscine.GetFontFile(fontType)

	output := piscine.Load(filename)
	if output == nil {
		fmt.Printf("Error loading font from file: %s\n", filename)
		return
	}
	piscine.PrintOutput(output, text, substring, *color)
}
