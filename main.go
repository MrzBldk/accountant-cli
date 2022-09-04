package main

import (
	"github.com/MrzBldk/accountant-cli/cmd"
	"github.com/MrzBldk/accountant-cli/database"
)

func main() {
	database.Init()
	cmd.Execute()
}
