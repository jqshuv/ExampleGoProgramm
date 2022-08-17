package main

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"github.com/tawesoft/golib/v2/dialog"
	"os"
)

func main() {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Printf("This is an example Program. \nMade by @jqhuv")
	} else {
		err := dialog.Info("This is an example Program. \nMade by @jqhuv")
		if err != nil {
			return
		}
	}

}
