package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cloudwatch.New(sess)

	resp, err := svc.DescribeAlarms(nil)
	if err != nil {
		fmt.Println("Got error getting alarm descriptions")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, alarm := range resp.MetricAlarms {
		fmt.Println(*alarm.AlarmName)
	}
}
