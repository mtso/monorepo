package main
// ami-ad5da9e9
// ami-86f4dbe6

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

	svc := ec2.New(sess, &aws.Config{Region: region})

	params := &ec2.RunInstancesInput{
		KeyName: aws.String("aws-eb3"),
		DryRun: aws.Bool(true),
		ImageId: aws.String("ami-86f4dbe6"),
		MinCount: aws.Int64(1),
		MaxCount: aws.Int64(1),
		InstanceType: aws.String("t1.micro"),
		SecurityGroups: []*string{ aws.String("sdk-connect") },
	}

	resp, err := svc.RunInstances(params)
	if aerr, ok := err.(awserr.Error); err != nil && !(ok && aerr.Code() == "DryRunOperation") {
		log.Fatal(err)
	}

	params.DryRun = aws.Bool(false)
	resp, err = svc.RunInstances(params)
	if err != nil {
		log.Fatal(err)
	}

	// var raw interface{}
	// err := json.Unmarshal(resp, &raw)

	// info := raw.(map[string]interface{})

	log.Printf("%+v\n", resp)

	// tagParams := &ec2.createTagsInput{
	// 	Resources: []*string{
	// 		aws.String(instanceId),
	// 	},
	// 	Tags: 
	// }
}
