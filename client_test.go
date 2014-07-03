package mapbox

import (
	"io/ioutil"
	"strings"
	"testing"
)

func readAccessToken(t *testing.T) (string, error) {
	contents, err := ioutil.ReadFile("MAPBOX_ACCESS_TOKEN")
	if err != nil {
		t.Fatalf("cannot read file MAPBOX_ACCESS_TOKEN; make sure to add this file and paste your Mapbox Access Token in there.\nError: %v", err)
		return "", nil
	}
	return strings.TrimSpace(string(contents)), nil
}

func TestMapboxDefaults(t *testing.T) {
	expected := "api.tiles.mapbox.com"
	if MapboxHost != expected {
		t.Errorf("expected mapbox host of %q, got: %q", expected, MapboxHost)
	}
	expected = "/v4"
	if MapboxPathPrefix != expected {
		t.Errorf("expected mapbox path prefix of %q, got: %q", expected, MapboxPathPrefix)
	}
}

func TestHTTPClient(t *testing.T) {
	c := NewClient("token")
	hc := c.HTTPClient()
	if hc == nil {
		t.Fatalf("expected HTTPClient() to never return nil, got: %v", hc)
	}
}

func TestBaseURL(t *testing.T) {
	c := NewClient("token")
	if c.HTTPS() {
		t.Error("expected HTTP scheme by default, got: HTTPS")
	}
	expected := "http://api.tiles.mapbox.com/v4"
	got := c.BaseURL()
	if got != expected {
		t.Errorf("expeced base URL of %q, got: %q", expected, got)
	}

	c = NewClient("token")
	c.SetHTTPS(true)
	if !c.HTTPS() {
		t.Error("expected HTTPS scheme, got: HTTP")
	}
	expected = "https://api.tiles.mapbox.com/v4"
	got = c.BaseURL()
	if got != expected {
		t.Errorf("expeced base URL of %q, got: %q", expected, got)
	}
}
