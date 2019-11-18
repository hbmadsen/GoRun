package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"

	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


func parseActivities(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.LazyQuotes = true

	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil

			}

			fmt.Println(err)

		}
		fmt.Println(row[0])

	}
}

func main() {

	parseActivities(os.Args[1])

	fmt.Println(rand.Intn(10))
	sess, err := session.NewSession()
	if err != nil {
		// Handle Session creation error
	}
	svc := s3.New(sess)
	fmt.Printf(svc.APIVersion)
}
