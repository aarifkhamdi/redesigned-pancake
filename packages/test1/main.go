package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var serverAddr string

func askForWelcome() {
	path := fmt.Sprintf("%s/api", serverAddr)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making request: %s", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body StatusCode=%d err=: %s", resp.StatusCode, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("StatusCode=%d Body: %s", resp.StatusCode, body)
		return
	}
	log.Printf("Status code: %d Message: %s", resp.StatusCode, body)
}

func Run() {
	addr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {
		panic("SERVER_ADDR not set")
	}
	serverAddr = addr
	time.AfterFunc(time.Second, askForWelcome)
}

func main() {
	Run()
}
