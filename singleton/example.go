package main

import (
	"log"
	"sync"
	"os"
	"fmt"
)

type hydraLogger struct{
	*log.Logger
	filename string
}

var hlogger *hydraLogger
var once sync.Once

func getInstance() *hydraLogger{
	once.Do(func(){
		fmt.Println("Create hydraLogger Once")
		hlogger = createInstance("hydralogger.log")
	})
	return hlogger
}

func createInstance(filename string) *hydraLogger{

	file , _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &hydraLogger{
		Logger: log.New(file, "Hydra ", log.Lshortfile),
		filename : filename,
	}
}

func main() {
	logger := getInstance()
	logger.Println("The first message")

	func(){
		logger := getInstance()
		logger.Println("This is the second message print from another func")
	}()
}