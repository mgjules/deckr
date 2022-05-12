package cmd

import "github.com/urfave/cli/v2"

// Commands is the list of CLI commands for the application.
var Commands = []*cli.Command{
	serve,
	version,
}
