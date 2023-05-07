package main

import (
	"io/ioutil"
	"math/rand"
	"time"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// go get github.com/gorilla/mux
)

// My server-- completed

func connect(w http.ResponseWriter, r *http.Request) {

	clientURL := r.RemoteAddr
	idt, err := json.Marshal(clientURL)
	if err != nil {
		fmt.Println("Error in the Connection")
	}

	res, ero := w.Write(idt)
	if ero != nil {
		fmt.Printf("Error : %s", ero)
	}
	fmt.Println(res)
	fmt.Println("Data Sent to Node", r.URL)

}

func givetime(w http.ResponseWriter, r *http.Request) {

	curr_time := time.Now()
	slot1 := curr_time.Add(time.Second * 5)

	idt, err := json.Marshal(slot1)
	if err != nil {
		fmt.Println("Error in the Connection")
	}

	_, ero := w.Write(idt)
	if ero != nil {
		fmt.Printf("Error : %s", ero)
	}
	fmt.Println("Data Sent to Node", r.URL)

}

func givedate(w http.ResponseWriter, r *http.Request) {

	curr_time := time.Now()
	slot1 := curr_time.Format("10-07-2008")

	idt, err := json.Marshal(slot1)
	if err != nil {
		fmt.Println("Error in the Connection")
	}

	_, ero := w.Write(idt)
	if ero != nil {
		fmt.Printf("Error : %s", ero)
	}
	fmt.Println("Data Sent to Node", r.URL)

}

func randnumber(w http.ResponseWriter, r *http.Request) {

	num := rand.Intn(50000) + 50000

	idt, err := json.Marshal(num)
	if err != nil {
		fmt.Println("Error in the Connection")
	}

	_, ero := w.Write(idt)
	if ero != nil {
		fmt.Printf("Error : %s", ero)
	}
	fmt.Println("Data Sent to Node", r.URL)

}

type clientInfo struct {
	Name string `json:"Name"`
	Age  int    `json:"age"`
}

func register(w http.ResponseWriter, r *http.Request) {

	registeration, _ := ioutil.ReadAll(r.Body)
	var cli clientInfo
	json.Unmarshal(registeration, &cli)

	fmt.Println(cli)

}

// IPC Socket for all requests (get/post)
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	//localhost:8082/connect
	myRouter.HandleFunc("/connect", connect).Methods("GET")
	myRouter.HandleFunc("/givetime", givetime).Methods("GET")
	myRouter.HandleFunc("/givedate", givedate).Methods("GET")
	myRouter.HandleFunc("/randnumber", randnumber).Methods("GET")
	myRouter.HandleFunc("/register", register).Methods("POST")

	// log.Fatal(http.ListenAndServe(":8082", myRouter))
	http.ListenAndServe(":8082", myRouter)
}

// main coroutine for the execution of the server
func main() {

	// Start Listner
	fmt.Println("Listening at 8082")
	fmt.Println("-----------------------")

	handleRequests()

}
