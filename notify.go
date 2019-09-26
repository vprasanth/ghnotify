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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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

	if err != nil {
		log.Fatal("Could not create request", err)
	}

	request.Header.Add("Authorization", token)

	resp, err := client.Do(request)

	if err != nil {
		log.Fatal("Error making request", err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	jsonBody := make([]map[string]interface{}, 10)

	err = decoder.Decode(&jsonBody)
	if err != nil {
		log.Fatal("Error decoding json", err)
	}

	out(len(jsonBody))

	//notifications := make(map[string]int)
	//for _, notification := range jsonBody {
	//	repoInfo := notification["repository"].(map[string]interface{})
	//	key := repoInfo["name"].(string)
	//	notifications[key]++
	//}

	//for key, count := range notifications {
	//	fmt.Printf("%s: %d | ", key, count)
	//}
	//fmt.Println("")

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
