package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
	"strings"
)

func nameViolation(bucketName, region string) string {
	if strings.HasPrefix(bucketName, region) {
		return false
	} else {
		return true
	}
}

func Handler(ctx context.Context, configEvent events.ConfigEvent) string {
	// Creating session via
	session := session.Must(session.NewSession())
	config := configservice.New(session, &aws.Config{})
	params := &configservice.PutEvaluationsInput{
		ResultToken: aws.String("String"), // Required
		Evaluations: []*configservice.Evaluation{
			{ // Required
				ComplianceResourceId:   aws.String("StringWithCharLimit256"), // Required
				ComplianceResourceType: aws.String("StringWithCharLimit256"), // Required
				ComplianceType:         aws.String("ComplianceType"),         // Required
				OrderingTimestamp:      aws.Time(time.Now()),                 // Required
				Annotation:             aws.String("StringWithCharLimit256"),
			},
			// More values...
		},
	}
	resp, err = config.PutEvaluations(params)
	evalution := evaluateCompliance(configurationItem, ruleParameters)
	if err != nil {
		// Print the error, cast err to the awserr.Error to get the Code and
		// Message from an error.
		fmt.PrinLn(err.Error())
		return
	}
	violation := nameViolation(bucketName, region)

}

func main() {
	lambda.Start(Handler)
}
