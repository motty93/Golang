package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://dark-sky.p.rapidapi.com/%7Blatitude%7D,%7Blongitude%7D?lang=en&units=auto"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	req.Header.Add("x-rapidapi-host", "dark-sky.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "ログインしてキーを取得")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(string(body))
}
