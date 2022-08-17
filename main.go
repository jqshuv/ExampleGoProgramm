package main

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"github.com/tawesoft/golib/v2/dialog"
	"os"
)

func main() {
	Run()
}

func Run() int {

	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Printf("This is an example Program. \nMade by @jqhuv")
		return 0
	} else {
		err := dialog.Info("This is an example Program. \nMade by @jqhuv")
		if err != nil {
			return 1 // exit if error
		}
		return 0 // exit if no error
	}
}
