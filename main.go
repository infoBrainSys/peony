package main

import (
	"peony/cmd"
	_ "peony/cmd"
	_ "peony/logic"
)

func main() {
	cmd.NewServer()
}
