package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Info struct {
	Items []struct {
		NewActiveUsers     int     `json:"new_active_users"`
		TotalUsers         int     `json:"total_users"`
		BadgesPerMinute    float64 `json:"badges_per_minute"`
		TotalBadges        int     `json:"total_badges"`
		TotalVotes         int     `json:"total_votes"`
		TotalComments      int     `json:"total_comments"`
		AnswersPerMinute   float64 `json:"answers_per_minute"`
		QuestionsPerMinute float64 `json:"questions_per_minute"`
		TotalAnswers       int     `json:"total_answers"`
		TotalAccepted      int     `json:"total_accepted"`
		TotalUnanswered    int     `json:"total_unanswered"`
		TotalQuestions     int     `json:"total_questions"`
		APIRevision        string  `json:"api_revision"`
	} `json:"items"`
	HasMore        bool `json:"has_more"`
	QuotaMax       int  `json:"quota_max"`
	QuotaRemaining int  `json:"quota_remaining"`
}

func httpExampleGetJsonString() {
	fmt.Println("--- GET json string ---")

	//	get http example
	resp, err := http.Get("https://api.stackexchange.com/2.2/info?site=stackoverflow")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//print json body
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))
}

func httpExampleGetJsonDecodeStruct() {
	fmt.Println("--- GET json to struct ---")

	// get http example
	resp, err := http.Get("https://api.stackexchange.com/2.2/info?site=stackoverflow")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var info Info

	// decode json to struct
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}

func httpExamplePost() {
	fmt.Println("--- POST json to struct ---")

	client := &http.Client{}
	jsonString := `{"id ":123456,"name":"Elvis","phones":["1234","2222","3333"]}`

	userJson, err := json.Marshal(jsonString)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://postman-echo.com/post", bytes.NewBuffer(userJson))

	// headers
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// decode json to map
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println(result["data"])

}

func main() {

	// http get
	httpExampleGetJsonString()
	httpExampleGetJsonDecodeStruct()

	// http post
	httpExamplePost()
}
