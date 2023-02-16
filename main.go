package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const Port_Number = "8090"
const Host_Address = "127.0.0.1"

type DividePageRequestBody struct {
	FirstOperator  float32 `json:"x"`
	SecondOperator float32 `json:"y"`
}

// HomePageHandler is the home page handler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	message := "This is the Home page"
	_, err := fmt.Println(message)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("x-rogery-header", "rogery-vargas-home-page")
	w.Write([]byte(message))
}

// AboutPageHandler is the about page handler
func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	mySum := addValue(2, 2)
	message := fmt.Sprintf("This is the About Page, and 2 + 2 is %d", mySum)
	_, err := fmt.Println(message)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("x-rogery-header", "rogery-vargas-about-page")
	w.Write([]byte(message))
}

// addValue add two values and return the sum
func addValue(x, y int) int {
	return x + y
}

func DividePageHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "an error occurred while trying to read the request's body")
		return
	}
	requestBody := DividePageRequestBody{}
	json.Unmarshal(bodyBytes, &requestBody)

	f, err := divideValues(requestBody.FirstOperator, requestBody.SecondOperator)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "The result of divide %f / %f is %f", requestBody.FirstOperator, requestBody.SecondOperator, f)
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/about", AboutPageHandler)
	http.HandleFunc("/divide", DividePageHandler)

	baseUrl := fmt.Sprintf("%v:%v", Host_Address, Port_Number)
	fmt.Printf("Trying to run the server on: http://%v\n", baseUrl)
	err := http.ListenAndServe(baseUrl, nil)
	if err != nil {
		log.Fatal("While trying to up the serve an error ocurred", err)
		os.Exit(1)
	}
}
