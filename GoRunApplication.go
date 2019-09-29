package main

import (
	"fmt"
	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		// Handle Session creation error
	}
	svc := s3.New(sess)
	fmt.Printf(svc.APIVersion)
}
