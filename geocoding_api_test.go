package mapbox

import (
	"strings"
	"testing"
)

func TestQueryByAddressURLs(t *testing.T) {
	accessToken := "pk.secretToken"

	tests := []struct {
		Request *QueryByAddressRequest
		URL     string
	}{
		{
			Request: &QueryByAddressRequest{
				Index: "mapbox.places-postcode-v1",
				Query: "20001",
			},
			URL: "http://api.tiles.mapbox.com/v4/geocode/mapbox.places-postcode-v1/20001.json?access_token=" + accessToken,
		},
		{
			Request: &QueryByAddressRequest{
				Index: "mapbox.places-province-v1",
				Query: "pennsylvania",
			},
			URL: "http://api.tiles.mapbox.com/v4/geocode/mapbox.places-province-v1/pennsylvania.json?access_token=" + accessToken,
		},
		{
			Request: &QueryByAddressRequest{
				Index: "mapbox.places-v1",
				Query: "1600 pennsylvania ave nw",
			},
			URL: "http://api.tiles.mapbox.com/v4/geocode/mapbox.places-v1/1600+pennsylvania+ave+nw.json?access_token=" + accessToken,
		},
	}

	geocoder := NewClient(accessToken).Geocoding()
	for _, test := range tests {
		got, err := geocoder.buildURL(test.Request)
		if err != nil {
			t.Fatalf("expeced no error, got: %v", err)
		}
		if got != test.URL {
			t.Errorf("expected %q, got: %q", test.URL, got)
		}
	}
}

func TestQueryByAddress(t *testing.T) {
	accessToken, err := readAccessToken(t)
	if err != nil {
		t.Fail()
		return
	}

	geocoder := NewClient(accessToken).Geocoding()
	req := &QueryByAddressRequest{
		Index: "mapbox.places-v1",
		Query: "Marienplatz 2,Munich,DE",
	}
	res, err := geocoder.QueryByAddress(req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if res == nil {
		t.Fatalf("expected result, got: %v", res)
	}
	if len(res.Features) < 1 {
		t.Fatalf("expected at least 1 feature, got: %v", len(res.Features))
	}
	f := res.Features[0]
	if len(f.Center) != 2 {
		t.Errorf("expected center with 2 coordinates, got: %v", len(f.Center))
	} else {
		if f.Center[0] != 11.541783 {
			t.Errorf("expected center longitude of 11.541783, got: %v", f.Center[0])
		}
		if f.Center[1] != 48.152471 {
			t.Errorf("expected center latitude of 48.152471 , got: %v", f.Center[1])
		}
	}
	if f.PlaceName != "Munich, Bayern, Germany" {
		t.Errorf("expected place name of %q, got: %v", "Munich, Bayern, Germany", f.PlaceName)
	}
	if f.Text != "Munich" {
		t.Errorf("expected text of %q, got: %v", "Munich", f.Text)
	}
	if f.ID != "city.676757" {
		t.Errorf("expected id %q, got: %v", "city.676757", f.ID)
	}
	if len(f.BoundingBox) != 4 {
		t.Errorf("expected bbox with 4 coordinates, got: %v", len(f.BoundingBox))
	}
}

func TestQueryByCity(t *testing.T) {
	accessToken, err := readAccessToken(t)
	if err != nil {
		t.Fail()
		return
	}

	geocoder := NewClient(accessToken).Geocoding()
	req := &QueryByAddressRequest{
		Index: "mapbox.places-v1",
		Query: "Munich",
	}
	res, err := geocoder.QueryByAddress(req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if res == nil {
		t.Fatalf("expected result, got: %v", res)
	}
	if len(res.Features) < 1 {
		t.Fatalf("expected at least 1 feature, got: %v", len(res.Features))
	}
	f := res.Features[0]
	if len(f.Center) != 2 {
		t.Errorf("expected center with 2 coordinates, got: %v", len(f.Center))
	} else {
		if f.Center[0] != 11.541783 {
			t.Errorf("expected center longitude of 11.541783, got: %v", f.Center[0])
		}
		if f.Center[1] != 48.152471 {
			t.Errorf("expected center latitude of 48.152471 , got: %v", f.Center[1])
		}
	}
	if !strings.Contains(f.PlaceName, req.Query) {
		t.Errorf("expected place name to contain %q, got: %v", req.Query, f.PlaceName)
	}
	if f.Text != "Munich" {
		t.Errorf("expected text of %q, got: %v", "Munich", f.Text)
	}
	if f.ID != "city.676757" {
		t.Errorf("expected id %q, got: %v", "city.676757", f.ID)
	}
	if len(f.BoundingBox) != 4 {
		t.Errorf("expected bbox with 4 coordinates, got: %v", len(f.BoundingBox))
	}
}
