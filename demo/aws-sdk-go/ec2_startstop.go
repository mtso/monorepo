package main

import (
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

	sess := session.Must(session.NewSession())
	region := os.Getenv("AWS_REGION")

	svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})

	switch strings.ToUpper(os.Args[1]) {
	case "STOP":
		params := &ec2.StopInstancesInput{
			InstanceIds: []*string{
				aws.String(os.Args[2]),
			},
			DryRun: aws.Bool(true),
		}
		resp, err := svc.StopInstances(params)
		if aerr, ok := err.(awserr.Error); err != nil && !(ok && aerr.Code() == "DryRunOperation") {
			log.Fatal(err.Error())
		}

		params.DryRun = aws.Bool(false)
		resp, err = svc.StopInstances(params)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("%+v\n", *resp)

	case "START":
		params := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String(os.Args[2]),
			},
			DryRun: aws.Bool(true),
		}
		_, err := svc.StartInstances(params)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "DryRunOperation" {
				params.DryRun = aws.Bool(false)

				resp, err := svc.StartInstances(params)
				if err != nil {
					log.Fatal(err)
				}

				log.Printf("%+v\n", *resp)
			} else {
				log.Fatal(err.Error())
			}
		} else {
			log.Fatal(err.Error())
		}

	default:
		break
	}
}
