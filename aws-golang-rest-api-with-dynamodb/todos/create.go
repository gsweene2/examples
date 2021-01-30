package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

    "fmt"
    "os"
	"encoding/json"
	"reflect"
)

type Item struct {
	Id       string  `json:"id,omitempty"`
    Title    string  `json:"title"`
    Details  string  `json:"details"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create DynamoDB client
    svc := dynamodb.New(sess)

	itemUuid := uuid.New().String()

	fmt.Println("Generted item uuid: %v", itemUuid)

	itemString := request.Body
	fmt.Println(reflect.TypeOf(itemString))
	itemStruct := Item{}
	fmt.Println(reflect.TypeOf(itemStruct))
	json.Unmarshal([]byte(itemString), &itemStruct)

    item := Item{
		Id: itemUuid,
        Title:   itemStruct.Title,
        Details:  itemStruct.Details,
    }

    av, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
        fmt.Println("Error marshalling item:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    tableName := os.Getenv("DYNAMODB_TABLE")

	fmt.Println("Putting item: %v", av)

    input := &dynamodb.PutItemInput{
        Item:      av,
        TableName: aws.String(tableName),
    }

    _, err = svc.PutItem(input)
    if err != nil {
        fmt.Println("Got error calling PutItem:")
        fmt.Println(err.Error())
        os.Exit(1)
	}
	
	item_marshalled, err := json.Marshal(item)

	fmt.Println("Returning item: ", string(item_marshalled))

	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{Body: string(item_marshalled), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}