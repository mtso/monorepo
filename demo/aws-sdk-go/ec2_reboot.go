package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess := session.Must(session.NewSession())
	region := aws.String(os.Getenv("AWS_REGION"))
	svc := ec2.New(sess, &aws.Config{
		Region: region,
	})

	params := &ec2.RebootInstancesInput{
		InstanceIds: []*string{
			aws.String(os.Args[1]),
		},
		DryRun: aws.Bool(true),
	}

	resp, err := svc.RebootInstances(params)
	if aerr, ok := err.(awserr.Error); err != nil && !(ok && aerr.Code() == "DryRunOperation") {
		log.Fatal(err)
	}

	params.DryRun = aws.Bool(false)
	resp, err = svc.RebootInstances(params)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("COMPLETE %+v\n", resp)
}
