package main

import (
	//"github.com/spf13/cobra"
	"github.com/yceman/cobra-go/cmd"
)

/*
	var templateFuncs = template.FuncMap{
		"trim": strings.TrimSpace,
	}

	type Command struct {
		PersistentPostRunE func(cmd *Command, args []string) error
	}
*/
func main() {
	cmd.Execute()
}
