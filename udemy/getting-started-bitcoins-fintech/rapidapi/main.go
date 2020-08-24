package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./config"
)

func main() {
	url := "https://community-open-weather-map.p.rapidapi.com/weather?id=2172797&units=%2522metric%2522%20or%20%2522imperial%2522&mode=xml%252C%20html&q=London%252Cuk"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", config.Config.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	fmt.Println(string(body))
}
