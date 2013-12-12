package geoip

import (
	"cgl.tideland.biz/asserts"
	"os"
	"testing"
)

func TestFindCity(t *testing.T) {
	assert := asserts.NewTestingAsserts(t, true)

	locator := NewLocator(os.Getenv("USER"), os.Getenv("KEY"))
        locator.Verbose = true
	g := locator.FindCity("186.52.156.53")

	assert.Equal(g.City.Names["en"], "Montevideo", "Wrong city name")
}
