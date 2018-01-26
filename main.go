package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func Handler() (Response, error) {
	result := hoge()
	return Response{
		Message: result,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

type ChidleyRoot314159 struct {
	Result *Result `xml:"result,omitempty" json:"result,omitempty"` // ZZmaxLength=0
}

type Address struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=75
}

type Coordinate struct {
	Lat     *Lat     `xml:"lat,omitempty" json:"lat,omitempty"`         // ZZmaxLength=0
	Lat_dms *Lat_dms `xml:"lat_dms,omitempty" json:"lat_dms,omitempty"` // ZZmaxLength=0
	Lng     *Lng     `xml:"lng,omitempty" json:"lng,omitempty"`         // ZZmaxLength=0
	Lng_dms *Lng_dms `xml:"lng_dms,omitempty" json:"lng_dms,omitempty"` // ZZmaxLength=0
}

type Google_maps struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=87
}

type Lat struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=9
}

type Lat_dms struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=12
}

type Lng struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=10
}

type Lng_dms struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=13
}

type Needs_to_verify struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=3
}

type Open_location_code struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=11
}

type Result struct {
	Address            *Address            `xml:"address,omitempty" json:"address,omitempty"`                       // ZZmaxLength=0
	Coordinate         *Coordinate         `xml:"coordinate,omitempty" json:"coordinate,omitempty"`                 // ZZmaxLength=0
	Google_maps        *Google_maps        `xml:"google_maps,omitempty" json:"google_maps,omitempty"`               // ZZmaxLength=0
	Needs_to_verify    *Needs_to_verify    `xml:"needs_to_verify,omitempty" json:"needs_to_verify,omitempty"`       // ZZmaxLength=0
	Open_location_code *Open_location_code `xml:"open_location_code,omitempty" json:"open_location_code,omitempty"` // ZZmaxLength=0
	Url                *Url                `xml:"url,omitempty" json:"url,omitempty"`                               // ZZmaxLength=0
	Version            *Version            `xml:"version,omitempty" json:"version,omitempty"`                       // ZZmaxLength=0
}

type Url struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=240
}

type Version struct {
	Text string `xml:",chardata" json:",omitempty"` // maxLength=3
}

func hoge() string {
	//arg := os.Args[1]
	arg := "東京都渋谷区恵比寿南1-16-12 ＡＢＣ・ＭＡＭＩＥＳ　３Ｆ"
	encodedAddress := url.QueryEscape(arg)
	url := fmt.Sprintf("http://www.geocoding.jp/api/?q=%s", encodedAddress)

	body, err := httpGet(url)
	fmt.Println(body)

	result := Result{}
	err = xml.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return ""
	}

	return fmt.Sprintf("Lat: %s, Lng: %s\n", result.Coordinate.Lat.Text, result.Coordinate.Lng.Text)
}

func httpGet(url string) (string, error) {
	response, _ := http.Get(url)
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}
