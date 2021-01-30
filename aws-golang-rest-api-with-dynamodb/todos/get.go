package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "fmt"
	"os"
	"encoding/json"
)

type Item struct {
	Id     string
    Title   string
    Details  string
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

	svc := dynamodb.New(sess)
	
	fetchingId := request.PathParameters["id"]
	
	fmt.Println("Derived fetchingId from path params: ", fetchingId)

    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
        Key: map[string]*dynamodb.AttributeValue{
            "Id": {
                S: aws.String(fetchingId),
            },
        },
    })

    if err != nil {
        fmt.Println(err.Error())
        return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
    }

    item := Item{}

    err = dynamodbattribute.UnmarshalMap(result.Item, &item)

    if err != nil {
        panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
    }

	fmt.Println("Found item: ", item)
	
	item_marshalled, err := json.Marshal(item)

	fmt.Println("Returning item: ", string(item_marshalled))

	return events.APIGatewayProxyResponse{Body: string(item_marshalled), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
