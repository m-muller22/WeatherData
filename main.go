package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var APIkey string

type WeatherData struct {
	Weather Weather `json:"weather"`
	Main    Main    `json:"main"`
	Name    string  `json:"name"`
}

type Weather []struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"01d"`
}

type Main struct {
	Temp       float32 `json:"temp"`
	Feels_like float32 `json:"feels_like"`
	Temp_min   float32 `json:"temp_min"`
	Temp_max   float32 `json:"temp_max"`
	Pressure   int     `json:"pressure"`
	Humidity   int     `json:"humidity"`
}

func main() {

	//read the APIKey for openweathermap
	content, err := os.ReadFile("apikey.txt")
	if err != nil {
		log.Fatalln(err)
	}
	APIkey = string(content)

	//read cities from cityList.txt and add to array string
	content, err = os.ReadFile("cityList.txt")
	if err != nil {
		log.Fatalln(err)
	}
	CityList := strings.Split(string(content), ",")

	var weatherD WeatherData
	var weatherDArray []WeatherData

	for i := 0; i < len(CityList); i++ {
		resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + CityList[i] + "&appid=" + APIkey + "&units=metric")
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(body, &weatherD)
		if err != nil {
			log.Fatalln(err)
		}
		weatherDArray = append(weatherDArray, weatherD)
	}

	fmt.Printf("----------\n")
	fmt.Printf("%+v\n", weatherDArray)
	fmt.Printf("----------\n")

}
