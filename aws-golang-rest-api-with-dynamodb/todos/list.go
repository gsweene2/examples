package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "fmt"
	"os"
	"encoding/json"
)

// type Item struct {
// 	Id       string  `json:"id,omitempty"`
//     Title    string  `json:"title"`
//     Details  string  `json:"details"`
// }

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

	svc := dynamodb.New(sess)
	
	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName:                 aws.String(os.Getenv("DYNAMODB_TABLE")),
	}

	// Make the DynamoDB Query API call
    result, err := svc.Scan(params)

    if err != nil {
        fmt.Println("Query API call failed:")
        fmt.Println((err.Error()))
        os.Exit(1)
    }

	fmt.Println("result: ", result)

	res, err := json.Marshal(result)
	if err != nil {
        fmt.Println("Got error marshalling result")
        fmt.Println(err.Error())
        return events.APIGatewayProxyResponse{Body: "Error getting items", StatusCode: 500}, nil
    }

	return events.APIGatewayProxyResponse{Body: string(res), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
