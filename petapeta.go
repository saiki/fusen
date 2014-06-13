package main

import (
	"fmt"
	"log"
//	"github.com/saiki/petapeta/wall"
	"./wall"
	"net/http"
	"os"
	"os/signal"
	"os/user"
)

var whiteboard *wall.Wall

const FILE_NAME = "exported"

func init() {
	whiteboard = new(wall.Wall).Init()
	path, err := getExportPath()
	if err != nil {
		panic(fmt.Sprintf("%v\n", err))
	}
	whiteboard.Import(path)
}

func main() {
	log.Println("start fusen...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("end fusen: in signal.Notify. captured:%v\n", sig)
			path, err := getExportPath()
			if err != nil {
				panic(fmt.Sprintf("%v\n", err))
			}
			err = whiteboard.Export(path)
			if err != nil {
				log.Fatalf("Export error: %v\n", err)
			}
			os.Exit(2)
		}
	}()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getExportPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir + string(os.PathSeparator) + FILE_NAME, nil
}
