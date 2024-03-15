package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	parse "github.com/plumium/inip/parse"
)

func main() {
	var fFlag = flag.String("f", "", "file path")
	flag.Parse()
	name := filepath.Base(*fFlag)
	content, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	input := string(content)
	ini := parse.Parse(name, input)
	b, err := json.MarshalIndent(ini, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
