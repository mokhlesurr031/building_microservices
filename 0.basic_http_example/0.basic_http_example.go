package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/helloworldstruct", helloWorldStructHandler)
	http.HandleFunc("/helloworld", helloWorldHandler)
	http.HandleFunc("/", rootHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloWorldStruct struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
	Author  string `json:"-"` //don't output this field'
	Date    string `json:"date,omitempty"`
	Id      int    `json:"id,string"` //convert output into string and rename to id
}

func helloWorldStructHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldStruct{Message: "Hello World", Number: 27}
	//log.Println(response)
	data, err := json.Marshal(response)
	//log.Println(string(data), err)
	if err != nil {
		panic("Oooops!")
	} else {
		fmt.Fprint(w, string(data))
	}
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Root\n")
}
