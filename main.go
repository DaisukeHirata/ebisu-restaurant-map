package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	geoJSON, err := proc("https://tabelog.com/tokyo/A1303/A130302/13020992/dtlmenu/lunch/")

	if err != nil {
		return events.APIGatewayProxyResponse{Body: `{"message": "JSON Marshal error"}`, StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{Body: geoJSON, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
	// geoJSON, _ := proc("https://tabelog.com/tokyo/A1303/A130302/13020992/dtlmenu/lunch/")
	// fmt.Println(geoJSON)
}

func proc(URL string) (string, error) {
	tabelogResult := GetAddressFromTabelogURL(URL)
	geoCoordResult := GetCoordinateFromAddress(tabelogResult.Address)
	return MarchallingToGeoJson(tabelogResult, geoCoordResult)
}
