package main

import (
	"fmt"
	"github.com/newrelic/go-agent"
	insights "github.com/newrelic/go-insights/client"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

type TestType struct {
	EventType    string `json:"eventType"`
	AwesomeScore int    `json:"AwesomeScore"`
}

func main() {
	config := newrelic.NewConfig("My Awesome Application", os.Getenv("NEW_RELIC_LICENSE_KEY"))
	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Error(err)
	}
	_ = os.Getenv("NEW_RELIC_ACCOUNT_ID")
	insightInsertKey := os.Getenv("NEW_RELIC_INSIGHTS_INSERT_KEY")

	client := insights.NewInsertClient(insightInsertKey, os.Getenv("NEW_RELIC_ACCOUNT_ID"))

	log.SetLevel(log.DebugLevel) // TODO

	if validationErr := client.Validate(); validationErr != nil {
		panic(validationErr)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		score := r.URL.Query().Get("score")
		i, err := strconv.Atoi(score)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Score must be a number, but %v\n", score)
		}

		testData := TestType{
			EventType:    "HelloEvent",
			AwesomeScore: i,
		}
		log.Debug(testData)

		if postErr := client.PostEvent(testData); postErr != nil {
			log.Errorf("Error: %v\n", postErr)
			w.WriteHeader(500)
		}
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/hello", handler))
	log.Fatal(http.ListenAndServe(":8001", nil))
}
