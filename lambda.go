package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/jmoiron/jsonq"
	"strings"
	"time"
)

// Structure for AWS Config Event, this is used to catch the config event and trigger the rule
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

func Handler(ctx context.Context, event ConfigEvent) (string, error) {

	session := session.Must(session.NewSession())
	config := configservice.New(session, &aws.Config{})
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(event.InvokingEvent))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	fmt.Print(aws.String(event.InvokingEvent))
	bucketName, err := jq.String("configurationItem", "resourceName")
	region, err := jq.String("configurationItem", "awsRegion")
	resourceType, err := jq.String("configurationItem", "resourceType")
	resourceID, err := jq.String("configurationItem", "resourceID")
	complianceValue := evaluateCompliance(bucketName, region)

	params := &configservice.PutEvaluationsInput{
		ResultToken: aws.String(event.ResultToken), // Required
		Evaluations: []*configservice.Evaluation{
			{
				ComplianceResourceId:   aws.String(resourceId), // Required
				ComplianceResourceType: aws.String(resourceType),
				ComplianceType:         aws.String(complianceValue), // Required
				OrderingTimestamp:      aws.Time(time.Now()),        // Required
			},
		},
	}
	resp, err := config.PutEvaluations(params)
	if err != nil {
		// Print the error, cast err to the awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
	}
	fmt.Println(resp)

	return bucketName, err

}

func main() {
	lambda.Start(Handler)
}
