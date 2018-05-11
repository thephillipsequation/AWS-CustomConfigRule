package main

import (
	"errors"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
	"strings"
	"encoding/json"
)

func nameViolation(bucketName, region string) string {
	if strings.HasPrefix(bucketName, region) {
		return false
	} else {
		return true
	}
}

func Handler(ctx context.Context, configEvent events.ConfigEvent ) string{
	config := configservice.New(sess, &aws.Config)
	invokingEvent :=
	configurationItem :=
	ruleParameters :=
	violation := nameViolation(bucketName, region)
	config := configservice.New(mySession)
	config.
	evalution := evaluateCompliance(configurationItem, ruleParameters)
	
}

func main() {
	lambda.Start(Handler)
}