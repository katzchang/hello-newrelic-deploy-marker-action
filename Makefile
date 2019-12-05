.PHONY: all test run

all: run

fmt:
	gofmt -w .

test:
	go test

run:
	go run main.go

revision:=$(shell git rev-parse HEAD)
description:=$(shell git log -1 --oneline | cut -b 9)
application_id:=475666074
user=$(shell git log -1 --pretty=format:'%an')
deploy-create:
	curl -X POST 'https://api.newrelic.com/v2/applications/$(application_id)/deployments.json' \
		 -H 'X-Api-Key:$(NEW_RELIC_REST_API_KEY)' -i \
		 -H 'Content-Type: application/json' \
		 -d \
		'{ \
		  "deployment": { \
			"revision": "$(revision)", \
			"changelog": "$(changelog)", \
			"description": "$(description)", \
			"user": "$(user)" \
		  } \
		}'
