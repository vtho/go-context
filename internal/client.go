package internal

import (
	"context"
	"log"
	"net/http"
	"time"
)

func httpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}
}

func getHTTPRequest() (*http.Request, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:9001/", nil)
	return request, cancel
}

func Client(w http.ResponseWriter, r *http.Request) {
	log.Println("[CLIENT] running now!")

	// custom http client to resue the TCP connection
	c := httpClient()

	// 1st request
	req, cancelContext := getHTTPRequest()
	defer cancelContext()
	_, err := c.Do(req)
	if err != nil {
		log.Printf("[CLIENT] %v", err)
	}

	// 2nd request
	req, cancelContext = getHTTPRequest()
	defer cancelContext()
	_, err = c.Do(req)
	if err != nil {
		log.Printf("[CLIENT] %v", err)
	}

	log.Println("[CLIENT] ran succesfully")
}
