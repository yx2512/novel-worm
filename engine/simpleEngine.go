package engine

import (
	"log"
)

type SimpleEngine struct{}

func (se SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)

		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}
