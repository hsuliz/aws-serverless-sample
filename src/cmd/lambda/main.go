package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"read-stats/internal"
)

func main() {
	lambda.Start(internal.Handler)
}
