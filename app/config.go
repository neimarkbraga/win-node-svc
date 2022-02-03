package app

import "github.com/neimarkbraga/win-node-svc/models"

// if you want to know advance options, descriptions can be found at https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/sc-create
// refer here to see other available options https://pkg.go.dev/golang.org/x/sys/windows/svc/mgr#Config

var Config = models.ServiceConfig{
	Name: "MyApp", // name of service. Space is not allowed here
	DisplayName: "My App", // display name of server. Spaces are allowed here
	Description: "Description for my app",

	// Node
	WorkingDirectory: "D:\\Users\\Makoy\\Projects\\MyApp", // root directory of node app
	EntryFile: "index.js", // node app entry file path (relative to WorkingDirectory)
	LogDirectory: "D:\\Users\\Makoy\\Projects\\MyApp\\logs", // directory where console outputs will be logged
}