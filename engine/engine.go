package engine

import (
	"justforunc/fetcher"
	"log"
)

func Run(sends ...Request)  {
	var requests []Request
	for _, value := range sends {
		requests = append(requests, value)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher error" +
				"fetching url is %s, %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		// ... imply expand all item of parseResult.Requests
		// put all item to requests
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}


	}

}
