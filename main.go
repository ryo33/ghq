package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var Version string = "0.7.2"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "ghq"
	app.Usage = "Manage GitHub repository clones"
	app.Version = Version
	app.Author = "motemen"
	app.Email = "motemen@gmail.com"
	app.Commands = Commands
	return app
}
