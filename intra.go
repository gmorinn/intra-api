package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron"
)

type UserResponse struct {
	Picture string `json:"picture"`
	Url     string `json:"url"`
	Title   string `json:"title"`
}

type NofifResponse struct {
	ID      string       `json:"id"`
	Title   string       `json:"title"`
	Class   string       `json:"class"`
	Content string       `json:"content"`
	Date    string       `json:"date"`
	User    UserResponse `json:"user"`
}

func getNotif() []*NofifResponse {

	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/guillaume.morin@epitech.eu/notification/message?format=json", os.Getenv("AUTH")), nil)
	if err != nil {
		fmt.Println("error 1")
		return nil
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error 2")
		return nil
	}
	//Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var resultRequest []*NofifResponse
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &resultRequest); err != nil {
		fmt.Println("error 3")
		return nil
	}
	return resultRequest
}

func main() {
	var cronTask *cron.Cron
	fmt.Println("start")
	http.HandleFunc("/", HelloHandler)
	fmt.Println("Server started at port 8080")
	runCron(&cronTask)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, there\n")
}
