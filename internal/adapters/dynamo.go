package adapters

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// InitializeDynamoClient creates the connection to the AWS and returns the DynamoDB client
func InitializeDynamoClient() *dynamodb.Client {
	// Load the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load AWS configuration: %v", err)
	}

	// Create the DynamoDB client using this configuration
	client := dynamodb.NewFromConfig(cfg)

	return client
}
