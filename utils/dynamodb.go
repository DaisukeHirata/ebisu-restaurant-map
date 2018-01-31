package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/satori/go.uuid"
)

func PutItem(tabelogResult TabelogResult, geocoordResult GeocoordResult) {

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
