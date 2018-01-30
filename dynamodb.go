package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/satori/go.uuid"
)

func PutItem() {

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
				S: aws.String("中国茶房8 恵比寿店 （チャイニーズカフェ・エイト）"),
			},
			"address": {
				S: aws.String("東京都渋谷区恵比寿南1-16-12 ＡＢＣ・ＭＡＭＩＥＳ　３Ｆ"),
			},
			"genre": {
				S: aws.String("中華料理 飲茶・点心 居酒屋"),
			},
			"url": {
				S: aws.String("https://tabelog.com/tokyo/A1303/A130302/13020992/dtlmenu/lunch/"),
			},
			"lat": {
				N: aws.String("35.644536"),
			},
			"lng": {
				N: aws.String("139.709527"),
			},
		},
	}

	putItem, putErr := svc.PutItem(putParams)
	if putErr != nil {
		panic(putErr)
	}
	fmt.Println(putItem)
}
