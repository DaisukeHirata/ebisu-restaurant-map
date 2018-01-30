package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	post := `OLが集うおしゃれな感じ。パン食べ放題です。
	https://tabelog.com/tokyo/A1303/A130302/13005718/`
	geoJSON, err := proc(post)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: `{"message": "JSON Marshal error"}`, StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{Body: geoJSON, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
	// post := `OLが集うおしゃれな感じ。パン食べ放題です。
	// https://tabelog.com/tokyo/A1303/A130302/13005718/`
	// geoJSON, _ := proc(post)
	// fmt.Println(geoJSON)
}

func proc(post string) (string, error) {
	URL := RegexTabelogURL(post)
	tabelogResult := GetAddressFromTabelogURL(URL)
	geoCoordResult := GetCoordinateFromAddress(tabelogResult.Address)
	return MarchallingToGeoJson(tabelogResult, geoCoordResult)
}

func HttpGet(url string) ([]byte, error) {
	response, _ := http.Get(url)
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
