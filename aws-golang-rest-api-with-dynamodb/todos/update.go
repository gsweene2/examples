package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "fmt"
	"os"
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

	svc := dynamodb.New(sess)
	
	fetchingId := request.PathParameters["id"]
	
	info := Item{
		Title: "Update",
		Details: "Success",
    }

    // expr, err := dynamodbattribute.MarshalMap(info)
    // if err != nil {
    //     fmt.Println("Got error marshalling info:")
    //     fmt.Println(err.Error())
    //     os.Exit(1)
    // }

    // key, err := dynamodbattribute.MarshalMap(item)
    // if err != nil {
    //     fmt.Println("Got error marshalling item:")
    //     fmt.Println(err.Error())
    //     os.Exit(1)
	// }
	
	fmt.Println("updating info.Title: ", info.Title)
	fmt.Println("updating info.Details: ", info.Details)


    // Update item in table Movies
    input := &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":t": {
                S: aws.String(info.Title),
			},
			":d": {
				S: aws.String(info.Details),
			},
        },
        TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
        Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(fetchingId),
			},
		},
        ReturnValues:              aws.String("UPDATED_NEW"),
        UpdateExpression:          aws.String("set Title = :t, Details = :d"),
    }

    _, err := svc.UpdateItem(input)
    if err != nil {
        fmt.Println(err.Error())
        return events.APIGatewayProxyResponse{Body: string("Yikes"), StatusCode: 500}, nil
    }

	return events.APIGatewayProxyResponse{Body: string("Done"), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
