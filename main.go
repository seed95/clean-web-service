package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/seed95/clean-web-service/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("cannot run the app, why? %v\n", aurora.Red(err))
	}
}
