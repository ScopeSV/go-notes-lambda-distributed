package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"example.com/notes/pkg/structs"
	"example.com/notes/pkg/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var note structs.NotePayload

	if err := json.Unmarshal([]byte(req.Body), &note); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}

	svc := utils.GetDynamoClient()

	_, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item: map[string]types.AttributeValue{
			"ID":      &types.AttributeValueMemberS{Value: utils.GenerateUUID()},
			"Title":   &types.AttributeValueMemberS{Value: note.Title},
			"Content": &types.AttributeValueMemberS{Value: note.Content},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(insert)
}
