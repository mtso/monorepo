package main

// ami-ad5da9e9
// ami-86f4dbe6
// ami-f4331a94

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Start script
const userdata = `#! /bin/bash
sudo echo "BEGIN USER DATA EXECUTION"
sudo netstat -tulnp
sudo pwd
sudo ls
sudo echo "HOME=$HOME"
sudo echo $HOME
sudo /home/ubuntu/echoserver &
sudo echo "END USER DATA EXECUTION"`

// i-0582a359d2b760659

func main() {
	sess := session.Must(session.NewSession())
	region := aws.String(os.Getenv("AWS_REGION"))

	svc := ec2.New(sess, &aws.Config{Region: region})
	startscript := base64.StdEncoding.EncodeToString([]byte(userdata))
	log.Println(startscript)

	params := &ec2.RunInstancesInput{
		KeyName:        aws.String("aws-eb2"),
		DryRun:         aws.Bool(true),
		ImageId:        aws.String("ami-1a1a337a"),
		MinCount:       aws.Int64(1),
		MaxCount:       aws.Int64(1),
		InstanceType:   aws.String("t2.micro"),
		SecurityGroups: []*string{aws.String("sdk-connect")},
		UserData:       aws.String(startscript),
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

	log.Printf("%+v\n", resp)
}
