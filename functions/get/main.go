package main

import (
	"fmt"

	"github.com/DaisukeHirata/ebisu-restaurant-map/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	geoJSON, _ := proc()
	return events.APIGatewayProxyResponse{Body: geoJSON, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
	//localDebug()
}

func localDebug() {
	geoJSON, _ := proc()
	fmt.Println(geoJSON)
}

func proc() (string, error) {
	restaurants := utils.ScanItems()

	geoJson, err := utils.MarchallingToGeoJson(restaurants)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return "", err
	}

	return geoJson, nil
}
