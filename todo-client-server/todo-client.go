package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/codegangsta/cli"
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
	// connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(-1)
	}

	// and setup a deferred function to close the connection properly before
	// we exit (just in case the connection hasn't been nil'ed out already)
	defer func() {
		conn.Close()
	}()

	// get our command from the command-line and send it to the server
	jb, _ := json.Marshal(cmdMap)
	sb := fmt.Sprintf("%s\n", string(jb))
	fmt.Printf("Command to send: %s", sb)
	// send to socket
	fmt.Fprintf(conn, sb)
	// listen for reply
	reply, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Reply from server: %s\n", reply)

}
