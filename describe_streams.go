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

	filter := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        aws.String("prod"),
		Descending:          aws.Bool(true),
		LogStreamNamePrefix: aws.String("bluemartini-website"),
	}

	resp, err := svc.DescribeLogStreams(filter)
	if err != nil {
		fmt.Println("Got error getting alarm descriptions")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(resp)
	for _, logStreams := range resp.LogStreams {
		fmt.Println(*logStreams.LogStreamName)
	}
}
