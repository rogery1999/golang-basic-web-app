package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const Port_Number = "8090"
const Host_Address = "127.0.0.1"

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/about", AboutPageHandler)

	baseUrl := fmt.Sprintf("%v:%v", Host_Address, Port_Number)
	fmt.Printf("Trying to run the server on: http://%v\n", baseUrl)
	err := http.ListenAndServe(baseUrl, nil)
	if err != nil {
		log.Fatal("While trying to up the serve an error ocurred", err)
		os.Exit(1)
	}
}
