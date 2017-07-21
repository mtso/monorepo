package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// AWS_ACCESS_KEY_ID
// AWS_SECRET_ACCESS_KEY
func main() {
	sess := session.Must(session.NewSession())
	region := os.Getenv("AWS_REGION")
	svc := ec2.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	params := &ec2.DescribeInstancesInput{}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		log.Println("error", region, err.Error())
		log.Fatal(err.Error())
	}
	log.Printf("%+v\n", *resp)
}
