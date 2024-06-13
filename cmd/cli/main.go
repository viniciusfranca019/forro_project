package main

import (
	"flag"
	"fmt"
	"forro_project/packages/Database/Migrate"
	"forro_project/packages/Database/Seeds"
	"forro_project/packages/Util"
)

func main() {
	Util.LoadEnv()
	command := flag.String("run", "help", "What command do you want to run")
	flag.Parse()

	switch *command {
	case "help":
		help()
	case "seeds":
		seeds()
	case "migrate":
		migrate()
	default:
		defaultCommand()
	}
}

func help() {
	fmt.Println("here are some commands we have")
	fmt.Println("help: display all commands and what they do")
	fmt.Println("seeds: Run all projects seeds")
	fmt.Println("migrate: Run all projects Migrations")
}

func seeds() {
	Seeds.Run()
}

func migrate() {
	Migrate.RunMigrate()
}

func defaultCommand() {
	fmt.Println("Command not implemented yet")
}
