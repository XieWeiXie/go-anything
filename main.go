package main

import (
	"go-anything/cmd"
	"log"
	"os"
)

func Execute() error {
	e := cmd.ROOT.Execute()
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}

func main() {
	e := Execute()
	if e != nil {
		os.Exit(1)
		return
	}

}
