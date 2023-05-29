package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

func sendRequest(task, url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Task", task)
	client := http.DefaultClient
	client.Timeout = time.Second * 15

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	answer := resp.Header.Values("Answer")
	if len(answer) == 0 {
		return "", errors.New("empty response header")
	}
	return answer[0], nil
}

func main() {

	set := os.Args
	if len(set) < 3 {
		fmt.Println("Not enough arguments, usage: cli substring url")
	}
	answer, err := sendRequest(set[1], set[2])
	if err != nil {
		fmt.Println("can't process task:", err.Error())
	}
	fmt.Println("Longest uniq substring: ", answer)

}
