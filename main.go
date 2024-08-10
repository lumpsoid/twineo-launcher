package main

import (
	"encoding/json"
	"fmt"
	"net/http"
    "bufio"
    "os"
    "strings"
)

// Struct to unmarshal JSON response
type Response struct {
	Data struct {
		Name string `json:"displayName"`
		Live bool   `json:"live"`
	} `json:"data"`
}

func main() {

    file, err := os.Open("subscription.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var usernames []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        username := strings.TrimSpace(scanner.Text())
        usernames = append(usernames, username)
    }
	// Check for any errors during scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Iterate over the array and make HTTP requests for each username
	for _, username := range usernames {
		url := fmt.Sprintf("https://twineo.exozy.me/api/user/%s", username)
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making GET request for %s: %s\n", username, err)
			continue
		}
		defer response.Body.Close()

        // Decode the JSON response
        var data Response
        err = json.NewDecoder(response.Body).Decode(&data)
        if err != nil {
            fmt.Println("Error decoding JSON response:", err)
            return
        }

        // Print information from the JSON response
        if data.Data.Live {
            fmt.Println(data.Data.Name, "live")
        } else {
            fmt.Println(data.Data.Name, "offline")
        }
	}
}
