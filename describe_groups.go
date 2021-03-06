package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cloudwatchlogs.New(sess)

	param := &cloudwatchlogs.DescribeLogGroupsInput{
		Limit:              aws.Int64(50),
		LogGroupNamePrefix: aws.String("prod"),
	}
	resp, err := svc.DescribeLogGroups(param)
	if err != nil {
		fmt.Println("Got error getting alarm descriptions")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(resp)
	for _, logGroups := range resp.LogGroups {
		fmt.Println(*logGroups.LogGroupName)
	}
}
