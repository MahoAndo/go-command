package main

import (
	"log"
	"os"
	"goCommand/cmd"
	"goCommand/conf"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//root initialization
	root, err := cmd.NewRoot()
	//Command root with cobra
	rootCmd := root.Build()

	builders := initBuilders()
	for _, b := range builders {
		command := b.Build()
		rootCmd.AddCommand(command)
	}

	err = rootCmd.Execute()

}

func initBuilders() []cmd.Builder {

	//load of setting
	c, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	var builders []cmd.Builder

	//add commands

	//conf
	confBuilder, err := cmd.NewConf(c)
	if err != nil {
		log.Fatal(err)
	}
	builders = append(builders, confBuilder)

	//userata
	userDataBuilder, err := cmd.NewUserData(c)
	if err != nil {
		log.Fatal(err)
	}
	builders = append(builders, userDataBuilder)

	//bookdata
	bookDataBuilder, err := cmd.NewBookData(c)
	if err != nil {
		log.Fatal(err)
	}
	builders = append(builders, bookDataBuilder)

	return builders
}
