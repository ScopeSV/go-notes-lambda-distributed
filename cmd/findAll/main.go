package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/notes/pkg/structs"
	"example.com/notes/pkg/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

func findAll() (events.APIGatewayProxyResponse, error) {
	svc := utils.GetDynamoClient()

	res, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprint("Error while retrieving data from DynamoDB", err),
		}, nil
	}

	var note []structs.Note

	if err = attributevalue.UnmarshalListOfMaps(res.Items, &note); err != nil {
		log.Fatalf("Error occured while umashalling, %v", err)
	}

	jsonRes, err := json.Marshal(note)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while decoding to string value",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonRes),
	}, nil
}

func main() {
	lambda.Start(findAll)
}
