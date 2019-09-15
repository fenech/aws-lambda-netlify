package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

type Event struct {
	Hook string `json:"hook"`
}

func HandleRequest(ctx context.Context, event Event) error {
	var hook string
	if event.Hook == "" {
		hook = os.Getenv("HOOK")
	} else {
		hook = event.Hook
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.netlify.com/build_hooks/"+hook, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
