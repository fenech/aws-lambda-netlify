package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Hook string `json:"hook"`
}

func HandleRequest(ctx context.Context, event Event) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.netlify.com/build_hooks/"+event.Hook, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func main() {
	lambda.Start(HandleRequest)
}
