package geoip

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const geoip_url = "https://geoip.maxmind.com/geoip/v2.0/"

type GeoIP struct {
	User string
	Key  string
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

type Geolocation struct {
	Country   Country   `json:"country"`
	Location  Location  `json:"location"`
	City      City      `json:"city"`
	Continent Continent `json:"continent"`
}

func NewLocator(user, key string) *GeoIP {
	return &GeoIP{user, key}
}

func (g *GeoIP) FindCity(ip string) Geolocation {
	var geolocation Geolocation

	client := &http.Client{}
	locator := geoip_url + "city/" + ip
	req, _ := http.NewRequest("GET", locator, nil)
	req.SetBasicAuth(g.User, g.Key)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &geolocation)

	return geolocation
}
