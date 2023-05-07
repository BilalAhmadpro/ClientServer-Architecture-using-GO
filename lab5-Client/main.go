package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

// My client-- completed

var address string = "http://localhost:8082"

func connect(_add string, _data chan string) {

	defer wg.Done()

	fmt.Println(_add)
	response, err := http.Get(_add + "/connect")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_data <- string(responseData)
}

func randomnumber(_add string, _randnum chan int) {

	defer wg.Done()

	response, err := http.Get(_add + "/randnumber ")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := binary.BigEndian.Uint64(responseData)
	_randnum <- int(data)
}

type clientInfo struct {
	Name string `json:"Name"`
	Age  int    `json:"age"`
}

func register(_add string, port int, _regs chan string) {

	defer wg.Done()

	v := _add + "/register"

	cli := clientInfo{
		Name: "Bilal Ahmad",
		Age:  22,
	}

	postBody, _ := json.Marshal(cli)

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(v, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	// _, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// sb := string(body)
	// _regs <- sb

}

func getdate(_add string, _data chan string) {

	defer wg.Done()

	response, err := http.Get(_add + "/givedate")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_data <- string(responseData)
}

func gettime(_add string, _data chan string) {

	defer wg.Done()

	response, err := http.Get(_add + "/givetime")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	_data <- string(responseData)
}

var dict map[string]interface{}
var sict map[string]interface{}
var peer map[string]interface{}
var wg sync.WaitGroup

func main() {

	_data := make(chan string)
	_randnum := make(chan int)
	_regs := make(chan string)
	_date := make(chan string)
	_time := make(chan string)
	wg.Add(5)

	//thread for the connection
	go connect(address, _data)
	data := <-_data

	fmt.Println("ip address:  " + data)

	//thread for the randomnumber
	go randomnumber(address, _randnum)

	data1 := <-_randnum

	fmt.Println("your random number is: ", data1)

	//thread for the register
	// imp note: result shown in server terminal
	go register(address, 7005, _regs)

	//thread for the getdate
	go getdate(address, _date)
	data3 := <-_date

	fmt.Println("current date is: ", data3)

	//thread for the gettime
	go gettime(address, _time)
	data4 := <-_time

	fmt.Println("your requested time is: ", data4)

	wg.Wait()

}
