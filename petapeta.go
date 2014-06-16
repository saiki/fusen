package main

import (
	"fmt"
	"log"
//	"github.com/saiki/petapeta/wall"
	"./model"
	"net/http"
	"os"
	"os/signal"
	"os/user"
)

var collection *model.Collection

const FILE_NAME = "exported"

func init() {
	collection = new(model.Collection)
	path, err := getExportPath()
	if err != nil {
		panic(fmt.Sprintf("%v\n", err))
	}
	collection.Import(path)
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
			err = collection.Export(path)
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
