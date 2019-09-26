package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {

	user, err := user.Current()
	check(err)

	path := filepath.Join(user.HomeDir, ".ghnotify")

	data, err := ioutil.ReadFile(path)
	check(err)

	token := fmt.Sprintf("Bearer %s", strings.TrimSpace(string(data)))
	notficationsEndpoint := "https://github.ibm.com/api/v3/notifications"
	client := http.Client{}

	request, err := http.NewRequest("GET", notficationsEndpoint, nil)
	check(err)

	request.Header.Add("Authorization", token)

	resp, err := client.Do(request)
	check(err)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Error fetching notifications. Response code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	jsonBody := make([]map[string]interface{}, 10)

	err = decoder.Decode(&jsonBody)
	check(err)

	out(len(jsonBody))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func out(l int) {

	icon := "🔔"

	if l > 10 && l < 20 {
		icon = "🤔"
	} else if l > 20 {
		icon = "😵"
	}

	fmt.Printf("%s %d\n", icon, l)
}
