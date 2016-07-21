package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JonahBraun/mctounicode/mctounicode"
)

func main() {
	scanner := mctounicode.NewMCScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "#") {
			continue
		}
		fmt.Println(s)
	}
}
