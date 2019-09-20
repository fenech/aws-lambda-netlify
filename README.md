# aws-lambda-netlify
AWS Lambda to trigger Netlify build, written in Go

## Build
```sh
GOOS=linux go build main.go
zip function.zip main
```

## Usage
The function sends a POST request to the following endpoint:
https://api.netlify.com/build_hooks/{hook}

The value of `hook` can come from:

1. (highest priority) an event of the form :
```json
{
    "hook": "YOUR_BUILD_HOOK"
}
```
2. the `HOOK` environment variable

Using the environment variable makes it possible to trigger the function using a CloudWatch scheduled event.
