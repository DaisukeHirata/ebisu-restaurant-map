package main

import (
	"fmt"
	"strings"

	"github.com/DaisukeHirata/ebisu-restaurant-map/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	result, err := proc(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: `{"message": "JSON Marshal error"}`, StatusCode: 500}, err
	}

	fmt.Println("Result: ", result)

	return events.APIGatewayProxyResponse{Body: result, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
	//localDebug()
}

func localDebug() {
	//	post := `OLが集うおしゃれな感じ。パン食べ放題です。
	//	https://tabelog.com/tokyo/A1303/A130302/13005718/`
	body := "token=fkYJB9Sl5SJkUikpIc38msW7&team_id=T0HSULB0F&team_domain=daisukihirata&service_id=306610673716&channel_id=C0HSQTX8V&channel_name=general&timestamp=1517273709.000317&user_id=U0HSULB39&user_name=dh&text=OL%E3%81%8C%E9%9B%86%E3%81%86%E3%81%8A%E3%81%97%E3%82%83%E3%82%8C%E3%81%AA%E6%84%9F%E3%81%98%E3%80%82%E3%83%91%E3%83%B3%E9%A3%9F%E3%81%B9%E6%94%BE%E9%A1%8C%E3%81%A7%E3%81%99%E3%80%82%0A%09%3Chttps%3A%2F%2Ftabelog.com%2Ftokyo%2FA1303%2FA130302%2F13005718%2F%3E"
	geoJSON, _ := proc(body)
	fmt.Println(geoJSON)
}

func proc(post string) (string, error) {
	decodedPost := utils.Unescape(post)
	textParam := textParam(decodedPost)
	URL := utils.RegexTabelogURL(textParam)
	if URL == "" {
		return "no URl", nil
	}
	message := messageInTextParam(textParam, URL)
	tabelogResult := utils.GetInfoFromTabelogURL(URL)
	geoCoordResult := utils.GetCoordinateFromAddress(tabelogResult.Address)
	utils.PutItem(tabelogResult, geoCoordResult, message)
	return "Success", nil
}

func textParam(post string) string {
	i := strings.Index(post, "text=")
	return strings.TrimSpace(post[i+5:])
}

func messageInTextParam(textParam string, URL string) string {
	message := strings.Replace(textParam, URL, "", -1)
	message = strings.Replace(message, "<", "", -1)
	message = strings.Replace(message, ">", "", -1)
	message = strings.TrimSpace(message)
	return message
}
