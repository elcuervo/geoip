package geoip

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const geoip_url = "https://geoip.maxmind.com/geoip/v2.0/"

type GeoIP struct {
	User    string
	Key     string
	Verbose bool
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	TimeZone  string  `json:"time_zone"`
}

type City struct {
	Names map[string]string `json:"names"`
}

type Country struct {
	Code  string            `json:"iso_code"`
	Names map[string]string `json:"names"`
}

type Continent struct {
	Code  string            `json:"code"`
	Names map[string]string `json:"names"`
}

type API struct {
	Remaining int `json:"queries_remaining"`
}

type Geolocation struct {
	Country   Country   `json:"country"`
	Location  Location  `json:"location"`
	City      City      `json:"city"`
	Continent Continent `json:"continent"`
	API       API       `json:"maxmind"`
}

func NewLocator(user, key string) *GeoIP {
	return &GeoIP{user, key, false}
}

func (g *GeoIP) check() {
	if g.User == "" || g.Key == "" {
		panic("You need a user and a key to use the service")
	}
}

func (g *GeoIP) FindCity(ip string) Geolocation {
	var geolocation Geolocation

	g.check()

	client := &http.Client{}
	locator := geoip_url + "city/" + ip
	if g.Verbose {
		log.Println(locator)
	}
	req, err := http.NewRequest("GET", locator, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(g.User, g.Key)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if g.Verbose {
		log.Println(string(body))
	}
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &geolocation)
	if err != nil {
		panic(err)
	}

	if g.Verbose {
		log.Println(geolocation.API.Remaining)
	}

	return geolocation
}
