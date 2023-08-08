package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter User Name : ")
	scanner.Scan()
	username := scanner.Text()

	// Call userDetails to get the list of follower if user exist
	userDetails(username)
}

func userDetails(username string) {

	request_url := fmt.Sprintf("https://api.github.com/users/%v", username)
	// fmt.Println("response ur = ", request_url)
	response, err := http.Get(request_url)

	// Check if Response have any error
	if err != nil {
		log.Fatal((err))
	}
	var data map[string]interface{}
	// Read response body data
	output, err := io.ReadAll(response.Body)

	json.Unmarshal([]byte(output), &data)

	value, exists := data["followers"]

	if !exists {
		fmt.Printf("User '%s' not found\n", username)
		return
	}

	fmt.Printf("Total no of follower for %s is %v \n", username, value)

	defer response.Body.Close()
}
