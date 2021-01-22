package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

type Message struct {
	ID     string `json:"_id"`
	Method string `json:"by"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Text   string `json:"body"`
}

func emailWorker(id int, messages <-chan Message, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for msg := range messages {
		time.Sleep(1 * time.Second)
		fmt.Println(fmt.Sprintf("msg: %s, method: %s, email: %s, group: %d", msg.Text, msg.Method, msg.Email, id))
	}
}

func phoneWorker(id int, messages <-chan Message, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for msg := range messages {
		time.Sleep(1 * time.Second)
		fmt.Println(fmt.Sprintf("msg: %s, method: %s, phone: %s, group: %d", msg.Text, msg.Method, msg.Phone, id))
	}
}

func main() {
	wg := new(sync.WaitGroup)
	byemail := make(chan Message)
	byphone := make(chan Message)
	messages := []Message{}

	path := os.Args[1]

	fileBytes, err := ioutil.ReadFile(path)

	checkError(err)

	json.Unmarshal(fileBytes, &messages)

	for i := 1; i <= 2; i++ {
		wg.Add(1)

		if i == 1 {
			go emailWorker(i, byemail, wg)
		} else {
			go phoneWorker(i, byphone, wg)
		}
	}

	for _, msg := range messages {
		switch msg.Method {
		case "email":
			byemail <- msg
		case "phone":
			byphone <- msg
		case "all":
			byemail <- msg
			byphone <- msg
		}
	}

	close(byemail)
	close(byphone)
	wg.Wait()
}
