package main

import (
	"github.com/newrelic/go-agent"
	insights "github.com/newrelic/go-insights/client"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"os"
)

type TestType struct {
	EventType    string `json:"eventType"`
	AwesomeScore int    `json:"AwesomeScore"`
}

func main() {
	config := newrelic.NewConfig("オーサムアプリケーション", os.Getenv("NEW_RELIC_LICENSE_KEY"))
	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Error(err)
	}
	_ = os.Getenv("NEW_RELIC_ACCOUNT_ID")
	insightInsertKey := os.Getenv("NEW_RELIC_INSIGHTS_INSERT_KEY")

	client := insights.NewInsertClient(insightInsertKey, "")

	log.SetLevel(log.DebugLevel) // TODO

	if validationErr := client.Validate(); validationErr != nil {
		//however it is appropriate to handle this in your use case
		log.Errorf("Validation Error!")
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		testData := TestType{
			EventType:    "HelloEvent",
			AwesomeScore: rand.Intn(100),
		}
		log.Debug(testData)

		if postErr := client.PostEvent(testData); postErr != nil {
			log.Errorf("Error: %v\n", postErr)
		}
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/hello", handler))
	log.Fatal(http.ListenAndServe(":8001", nil))
}
