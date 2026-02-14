package repository

import (
	"context"
	"net/http"
	"strconv"

	"github.com/JoaoVitor615/URL-shortener/internal/domain"
	"github.com/JoaoVitor615/URL-shortener/internal/pkg/apperrors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoRepository struct {
	Client    *dynamodb.Client
	TableName string
}

var (
	ErrURLNotFound  = apperrors.New("URL not found", http.StatusNotFound)
	ErrUnmarshalURL = apperrors.NewWithErr("failed to unmarshal url", http.StatusInternalServerError)
	ErrMarshalURL   = apperrors.NewWithErr("failed to marshal url", http.StatusInternalServerError)
	ErrSaveURL      = apperrors.NewWithErr("failed to save url", http.StatusInternalServerError)
	ErrGetURL       = apperrors.NewWithErr("failed to get url", http.StatusInternalServerError)
)

func NewDynamoRepository(client *dynamodb.Client, tableName string) *DynamoRepository {
	return &DynamoRepository{
		Client:    client,
		TableName: tableName,
	}
}

func (s *DynamoRepository) GetLongURL(ctx context.Context, longURL string) (url *domain.URL[int], err error) {
	input := &dynamodb.ScanInput{
		TableName:        aws.String(s.TableName),
		FilterExpression: aws.String("long_url = :url"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":url": &types.AttributeValueMemberS{Value: longURL},
		},
	}

	output, err := s.Client.Scan(ctx, input)
	if err != nil {
		return nil, ErrGetURL(err)
	}

	if len(output.Items) == 0 {
		return nil, ErrURLNotFound
	}

	var result domain.URL[int]
	err = attributevalue.UnmarshalMap(output.Items[0], &result)
	if err != nil {
		return nil, ErrUnmarshalURL(err)
	}

	return &result, nil
}

func (s *DynamoRepository) SaveURL(ctx context.Context, url *domain.URL[int]) error {
	item, err := attributevalue.MarshalMap(url)
	if err != nil {
		return ErrMarshalURL(err)
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(s.TableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}

	_, err = s.Client.PutItem(ctx, input)
	if err != nil {
		return ErrSaveURL(err)
	}
	return nil
}

func (s *DynamoRepository) GetURL(ctx context.Context, id int) (*domain.URL[int], error) {
	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberN{Value: strconv.Itoa(id)},
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(s.TableName),
		Key:       key,
	}

	output, err := s.Client.GetItem(ctx, input)
	if err != nil {
		return nil, ErrGetURL(err)
	}

	if output.Item == nil {
		return nil, ErrURLNotFound
	}

	var url domain.URL[int]
	err = attributevalue.UnmarshalMap(output.Item, &url)
	if err != nil {
		return nil, ErrUnmarshalURL(err)
	}

	return &url, nil
}
