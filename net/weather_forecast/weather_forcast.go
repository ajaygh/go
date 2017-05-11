package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type DataPoint struct {
	Time                   float64
	Summary                string
	Icon                   string
	SunriseTime            float64
	SunsetTime             float64
	PrecipIntensity        float64
	PrecipIntensityMax     float64
	PrecipIntensityMaxTime float64
	PrecipProbability      float64
	PrecipType             string
	PrecipAccumulation     float64
	Temperature            float64
	TemperatureMin         float64
	TemperatureMinTime     float64
	TemperatureMax         float64
	TemperatureMaxTime     float64
	DewPoint               float64
	WindSpeed              float64
	WindBearing            float64
	CloudCover             float64
	Humidity               float64
	Visibility             float64
	Ozone                  float64
}

type Forecast struct {
	Latitude  float64
	Longitude float64
	Timezone  string
	Offset    float64
	Currently DataPoint
	Junk      string
}

func main() {
	//Create a waitgroup to manage the goroutines
	var wg sync.WaitGroup

	// Perform 4 concurrent queries
	wg.Add(4)
	for query := 0; query < 4; query++ {
		go Get(query, &wg)
	}

	//Wait for all the queries to complete
	wg.Wait()
	fmt.Println("All queries completed")
}

func Get(query int, wg *sync.WaitGroup){
	//Decrement the counter to show its completion
	defer wg.Done()

	addr := []{"Delhi, India", "Moscow, Russia"}

	//Geocoding api
	//QueryEscape escapes the address string so 
	//it can be safely placed inside a URL query
	safeAddr := url.QueryEscape(addr[query])
	fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=%s", safeAddr)
	
	//Build the request
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		panic(err)
	}

	//For control over HTTP client, redirect policy
	// and other settings create a client
	client := &http.Client{}
	//Send the request via the client
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	//Callers should close response.Body when Done
	defer resp.Body.Close()

	//Read the content into a byte array
	body , err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err2)
	}

	var res interface{}
	json.Unmarshal(body, &res)

	//Get lat, long
	 lat, _ := res["results"][0]["geometry"]["location"]["lat"]
	 lng, _ := res["results"][0]["geometry"]["location"]["lng"]

	 //Forecast api
	 url := Sprintf("https://api.darksky.net/forecast/myapikey/%.13f,%.13f?units=ca",lat, lng)
	 resp, err := http.get(url)
	if err != nil {
		log.Fatal(err)
	}

	fbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fata(err)
	}

	var f Forecast
	json.Unmarshal(fbody, &f)

	//Print detail results
	fmt.Println("The weather at ", addr[query])
	fmt.Println("Timezone at ", f.Timezone)
	fmt.Println("Temperature in celsius = ", f.Currently.Temperature)
	fmt.Println("Summary ", f.Currently.Summary)
}

