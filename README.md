# aws-lambda-netlify
AWS Lambda to trigger Netlify build, written in Go

## Build
```sh
GOOS=linux go build main.go
zip function.zip main
```

## Usage
Send an event of the form:
```json
{
    "hook": "YOUR_BUILD_HOOK"
}
```
