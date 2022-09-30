package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/alvii147/FunFaker/data"
	"github.com/alvii147/FunFaker/router"
	"github.com/alvii147/FunFaker/utils"
)

// func type that takes in command-line arguments
type CommandFunc func(args []string)

// command description and function
type CommandInfo struct {
	description string
	function    CommandFunc
}

// map of commands
type CommandMap map[string]CommandInfo

// define map of available commands
var COMMANDS = CommandMap{
	"validate": {
		description: "parse and validate data",
		function:    validate,
	},
	"server": {
		description: "run HTTP server",
		function:    server,
	},
}

// run validation tests on data files
func validate(args []string) {
	// parse command-line flags
	flagSet := flag.NewFlagSet("datahealth", flag.ExitOnError)
	autoFixPtr := flagSet.Bool("autofix", false, "Fix issues automatically where possible")

	flagSet.Parse(args)
	autofix := *autoFixPtr

	err := data.Validate(autofix)
	if err != nil {
		fmt.Println("Validation failed:", err)
		fmt.Println("Run with flag -autofix to fix issues automatically where possible")
	} else {
		fmt.Println("Validation passed")
	}
}

// run server
func server(args []string) {
	// parse command-line flags
	flagSet := flag.NewFlagSet("server", flag.ExitOnError)
	hostnamePtr := flagSet.String("hostname", "", "Server host name")
	portPtr := flagSet.Int64("port", 8080, "Server port number")

	flagSet.Parse(args)
	hostname := *hostnamePtr
	port := strconv.FormatInt(*portPtr, 10)

	router.Routing()

	addr := hostname + ":" + port
	utils.LogInfo("Server running on " + addr)
	http.ListenAndServe(addr, nil)
}

// print program usage
func usage() {
	fmt.Println("\nUsage:")
	fmt.Println("\tgo run main.go <command> [options]")
	fmt.Println("\nCommands:")
	for name, info := range COMMANDS {
		fmt.Printf("\t%s:\t%25s\n", name, info.description)
	}
}

func main() {
	// print usage and exit if no command executed
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	commandInfo, ok := COMMANDS[os.Args[1]]
	// print usage and exit if invalid command executed
	if !ok {
		usage()
		os.Exit(127)
	}

	// execute command
	commandInfo.function(os.Args[2:])
}
