package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/satori/go.uuid"
)

type Restaurant struct {
	ID        string  `dynamodb:"id"`
	CreatedAt int     `dynamodb:"createdAt"`
	Address   string  `dynamodb:"address"`
	Genre     string  `dynamodb:"genre"`
	Message   string  `dynamodb:"message"`
	Lat       float64 `dynamodb:"lat"`
	Lng       float64 `dynamodb:"lng"`
	Name      string  `dynamodb:"name"`
	URL       string  `dynamodb:"url"`
}

func PutItem(tabelogResult TabelogResult, geocoordResult GeocoordResult, message string) {

	// session
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)
	uuid := uuid.Must(uuid.NewV4()).String()
	createdAt := strconv.FormatInt(time.Now().Unix(), 10)

	// PutItem
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("ebisu-restaurant"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(uuid),
			},
			"createdAt": {
				N: aws.String(createdAt),
			},
			"name": {
				S: aws.String(tabelogResult.Name),
			},
			"address": {
				S: aws.String(tabelogResult.Address),
			},
			"genre": {
				S: aws.String(tabelogResult.Genre),
			},
			"message": {
				S: aws.String(message),
			},
			"url": {
				S: aws.String(tabelogResult.URL),
			},
			"lat": {
				N: aws.String(geocoordResult.Coordinate.Lat.Text),
			},
			"lng": {
				N: aws.String(geocoordResult.Coordinate.Lng.Text),
			},
		},
	}

	putItem, putErr := svc.PutItem(putParams)
	if putErr != nil {
		panic(putErr)
	}
	fmt.Println(putItem)
}

func ScanItems() []Restaurant {
	// session
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("ebisu-restaurant"),
	}

	result, err := svc.Scan(params)
	if err != nil {
		fmt.Errorf("failed to make Query API call, %v", err)
	}

	obj := []Restaurant{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &obj)
	if err != nil {
		fmt.Errorf("failed to unmarshal Query result items, %v", err)
	}

	return obj
}
