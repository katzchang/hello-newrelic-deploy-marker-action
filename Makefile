.PHONY: all test run

all: run

fmt:
	gofmt -w .

test:
	go test

run:
	go run main.go

revision:=$(shell git rev-parse HEAD)
description:=$(shell git log -1 --oneline)
application_id:=475666074
user=$(shell git log -1 --pretty=format:'%an')
deploy-create:
	curl -X POST 'https://https://github.com/katzchang/hello-newrelic-deploy-marker-action/commit/44cf1eda53423295ecee4a5aa8bc628a95d27aeb/checks?check_suite_id=343024029/v2/applications/$(application_id)/deployments.json' \
		 -H 'X-Api-Key:$(NEW_RELIC_REST_API_KEY)' -i \
		 -H 'Content-Type: application/json' \
		 -d '{"deployment":{"revision":"$(revision)","changelog":"$(changelog)","description":"$(description)","user": "$(user)"}}'
