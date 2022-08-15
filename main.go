package main

import "github.com/tawesoft/golib/v2/dialog"

func main() {
	err := dialog.Info("This is an example Program. \n Made by @jqhuv")
	if err != nil {
		return
	}
}
