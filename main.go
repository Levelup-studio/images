package main

import (
	"github.com/Levelup-studio/images/cmd"
)

//go:generate pkger -include /static -o build

func main() {
	cmd.Execute()
}
