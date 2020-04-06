package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type country struct {
	Country  string `json:"Country"`
	Province string `json:"Province"`
	Lat      string `json:"Lat"`
	Lon      string `json:"Lon"`
	Date     string `json:"Date"`
	Cases    int    `json:"Cases"`
	Status   string `json:"Status"`
}

type data struct {
	Country country
	Status  string
	Err     error
}

func getStatus(c string, s string, countryChan chan *data) {
	r, err := http.Get("https://api.covid19api.com/total/country/" + c + "/status/" + s)
	if err != nil {
		countryChan <- &data{country{}, s, err}
	}
	defer r.Body.Close()

	countries := make([]country, 0)
	err = json.NewDecoder(r.Body).Decode(&countries)
	if err != nil {
		countryChan <- &data{country{}, s, err}
	}

	if len(countries) == 0 {
		err := errors.New("no results returned for country \"" + c + "\"")
		countryChan <- &data{country{}, s, err}
	}

	country := countries[len(countries)-1]

	countryChan <- &data{country, s, err}
}

func getStatuses(country string, statuses []string) (statusMap map[string]int, err error) {
	countryChan := make(chan *data)

	for _, status := range statuses {
		go getStatus(country, status, countryChan)
	}

	statusMap = make(map[string]int)
	for i := 0; i < len(statuses); i++ {
		d := <-countryChan
		if d.Err != nil {
			return statusMap, d.Err
		}
		statusMap[d.Status] = d.Country.Cases
	}
	close(countryChan)

	// calculate active
	statusMap["active"] = statusMap["confirmed"] - (statusMap["recovered"] + statusMap["deaths"])

	return statusMap, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errors.New("country name required; try us, canada or south-africa"))
	}
	country := os.Args[1]

	statuses := []string{"confirmed", "recovered", "deaths"}
	statusMap, err := getStatuses(country, statuses)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range statusMap {
		fmt.Printf("%-10s-%9d\n", k, v)
	}
}
