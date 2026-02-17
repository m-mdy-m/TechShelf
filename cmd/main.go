package main

import (
	"os"

	"github.com/m-mdy-m/TechShelf/internal/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
