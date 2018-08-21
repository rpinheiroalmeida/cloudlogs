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

	// filter := &cloudwatchlogs.DescribeLogStreamsInput{
	// 	LogGroupName:        aws.String("prod"),
	// 	Descending:          aws.Bool(true),
	// 	LogStreamNamePrefix: aws.String("bluemartini-website"),
	// }

	filter := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String("prod"),
		LogStreamName: aws.String("bluemartini-website/bluemartini-website/fffdcf1f-7cb8-44a0-9caa-1daeffd8ad6b"),
	}

	resp, err := svc.GetLogEvents(filter)
	if err != nil {
		fmt.Println("Got error getting alarm descriptions")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(resp)
	// for _, logStreams := range resp.LogStreams {
	// 	fmt.Println(*logStreams.LogStreamName)
	// }
}
