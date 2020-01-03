package main

import (
	"bufio"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	usernameList, err := readLines("usernames.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	file, err := os.Create("available.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Enter function number (1..3)")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(char), "has been pressed...")

	switch char {

	case '1':
		for _, line := range usernameList {
			time.Sleep(500 * time.Millisecond)
			if checkUsername1(line) == true {
				fmt.Println(line, "IS AVAILABLE TO REGISTER!!!")
				file.WriteString(line + "\r\n")
			} else {
				fmt.Println(line, "is not available to register...")
			}
		}
		break
	case '2':
		for _, line := range usernameList {
			time.Sleep(500 * time.Millisecond)
			if checkUsername2(line) == true {
				fmt.Println(line, "IS AVAILABLE TO REGISTER!!!")
				file.WriteString(line + "\r\n")
			} else {
				fmt.Println(line, "is not available to register...")
			}
		}
		break
	case '3':
		for _, line := range usernameList {
			time.Sleep(500 * time.Millisecond)
			if checkUsername3(line) == true {
				fmt.Println(line, "IS AVAILABLE TO REGISTER!!!")
				file.WriteString(line + "\r\n")
			} else {
				fmt.Println(line, "is not available to register...")
			}
		}
	}

	file.Close()

	fmt.Println("\nPress 'Enter' to close...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	//fmt.Println(checkUsername1(line))
	//fmt.Println(checkUsername2(line))
	//fmt.Println(checkUsername3(line))

}

func checkUsername1(username string) bool {
	resp, err := http.Get("https://mixer.com/api/v1/channels?scope=names&limit=1&q=" + username)

	if err != nil {
		// handle error
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	resp.Body.Close()

	usernameStatus := gjson.Get(bodyString, "0.user.username").String()

	if strings.ToLower(username) == strings.ToLower(usernameStatus) {
		return false // already registered
	} else {
		return true // available to register
	}

}

// alternate ways of checking username availability 

func checkUsername2(username string) bool {
	resp, err := http.Get("https://mixer.com/api/v1/channels/" + username)

	if err != nil {
		// handle error
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	resp.Body.Close()

	usernameStatus := gjson.Get(bodyString, "token").String()

	if strings.ToLower(username) == strings.ToLower(usernameStatus) {
		return false // already registered
	} else {
		return true // available to register
	}

}

func checkUsername3(username string) bool {
	resp, err := http.Get("https://mixer.com/api/v1/channels/" + username)

	if err != nil {
		// handle error
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	resp.Body.Close()

	usernameStatus := gjson.Get(bodyString, "message").String()

	if "Channel not found." != usernameStatus {
		return false // already registered
	} else {
		return true // available to register
	}

}

func readLines(path string) ([]string, error) {
	// https://stackoverflow.com/a/18479916
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines, scanner.Err()
}
