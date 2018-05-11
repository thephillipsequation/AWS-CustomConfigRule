package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"log"
	"strings"
	"time"
)

type ConfigEvent struct {
	AccountID        string `json:"accountId"`     // The ID of the AWS account that owns the rule
	ConfigRuleArn    string `json:"configRuleArn"` // The ARN that AWS Config assigned to the rule
	ConfigRuleID     string `json:"configRuleId"`
	ConfigRuleName   string `json:"configRuleName"` // The name that you assigned to the rule that caused AWS Config to publish the event
	EventLeftScope   bool   `json:"eventLeftScope"` // A boolean value that indicates whether the AWS resource to be evaluated has been removed from the rule's scope
	ExecutionRoleArn string `json:"executionRoleArn"`
	InvokingEvent    string `json:"invokingEvent"`  // If the event is published in response to a resource configuration change, this value contains a JSON configuration item
	ResultToken      string `json:"resultToken"`    // A token that the function must pass to AWS Config with the PutEvaluations call
	RuleParameters   string `json:"ruleParameters"` // Key/value pairs that the function processes as part of its evaluation logic
	Version          string `json:"version"`
}

func evaluateCompliance(bucketName, region string) string {
	if strings.HasPrefix(bucketName, region) {
		return "COMPLIANT"
	} else {
		return "NON_COMPLIANT"
	}
}

func Handler(ctx context.Context, event ConfigEvent) string {
	// Creating session via
	session := session.Must(session.NewSession())
	config := configservice.New(session, &aws.Config{})
	fmt.Println(configEvent)
	// params := &configservice.PutEvaluationsInput{
	// 	ResultToken: aws.String("String"), // Required
	// 	Evaluations: []*configservice.Evaluation{
	// 		{ // Required
	// 			ComplianceResourceId:   aws.String(resourceId), // Required
	// 			ComplianceResourceType: aws.String(resourceType), // Required
	// 			ComplianceType:         aws.String(evaluateCompliance()),         // Required
	// 			OrderingTimestamp:      aws.Time(time.Now()),                 // Required
	// 		},
	// 		// More values...
	// 	},
	// }
	// resp, err = config.PutEvaluations(params)
	if err != nil {
		// Print the error, cast err to the awserr.Error to get the Code and
		// Message from an error.
		fmt.PrinLn(err.Error())
		return
	}
	// violation := nameViolation(bucketName, region)

}

func main() {
	lambda.Start(Handler)
}
