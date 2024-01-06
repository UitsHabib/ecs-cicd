package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if err := serveRest(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
