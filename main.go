package main

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-plugin-template/commands"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "Upload Support Bundle"
	app.Description = "Uploading a support bundle to JFrog SaaS instance (https://supportlogs.jfrog.com/logs/) using CLI command"
	app.Version = "v0.0.1"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetHelloCommand()}
}

//  curl -i -T {filename.ext} “https://supportlogs.jfrog.com/logs/<Case-Number>/”
