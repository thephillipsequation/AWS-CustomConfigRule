package main

import (
		"fmt"
		"context"
		"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
		Name string `"json:name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
		return fmt.Sprintdf("Hello  ya fool %s!", name.Name), nil
}

func main() {
		lambda.Start(HandleRequest)
}