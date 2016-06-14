package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/tjmcs/go-api-examples/todo-app/lib"
)

func main() {

	// setup our CLI
	var (
		cmdFlag = cli.StringFlag{
			Name:  "command, c",
			Usage: "Command to execute: add, delete, list, or details",
		}
		itemFlag = cli.StringFlag{
			Name:  "todo-details, d",
			Usage: "ToDo Item Details: JSON_STR",
		}
		idFlag = cli.StringFlag{
			Name:  "todo-id, i",
			Usage: "ToDo Item ID",
		}
	)
	app := cli.NewApp()
	app.Name = "todo-client"
	app.Version = "1.0.0"
	app.Usage = "ToDo command-line client"
	app.Flags = []cli.Flag{
		cmdFlag,
		itemFlag,
		idFlag,
	}
	app.Action = action
	app.Run(os.Args)
}

func action(ctx *cli.Context) {

	// build command map from input flags
	cmdMap := make(map[string]string)
	if ctx.IsSet("command") {
		cmdMap["command"] = ctx.String("command")
	}
	if ctx.IsSet("todo-details") {
		cmdMap["todo-details"] = ctx.String("todo-details")
	}
	if ctx.IsSet("todo-id") {
		cmdMap["todo-id"] = ctx.String("todo-id")
	}

	// call the underlying function to process the input command
	reply := lib.ProcessCmd(cmdMap)
	// and print the reply
	fmt.Printf("Reply from server: %s\n", reply)

}
