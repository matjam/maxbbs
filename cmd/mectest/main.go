package main

import (
	"fmt"
	"os"

	"github.com/matjam/mecca"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mectest <template>")
		os.Exit(1)
	}

	interpreter := mecca.NewInterpreter(mecca.WithTemplateRoot("cmd/mectest"))
	err := interpreter.RenderTemplate(os.Args[1], map[string]any{"bbsversion": "v1.0.0", "bbsname": "MaxBBS", "sysopname": "Max"})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
