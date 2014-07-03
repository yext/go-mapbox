package mapbox

import (
	_ "bytes"
	_ "image/png"
	_ "io/ioutil"
	"testing"
)

func TestGetImageURLs(t *testing.T) {
	accessToken := "pk.secretToken"
	mapId := "olivere.167ai10"

	tests := []struct {
		Request *ImageRequest
		URL     string
	}{
		{
			Request: &ImageRequest{
				Longitude: 11.54165,
				Latitude:  48.151313,
				Zoom:      9,
				Width:     500,
				Height:    300,
				Format:    "png",
			},
			URL: "http://api.tiles.mapbox.com/v4/" + mapId + "/11.541650,48.151313,9/500x300.png?access_token=" + accessToken,
		},
		{
			Request: &ImageRequest{
				Longitude: 11.54165,
				Latitude:  48.151313,
				Zoom:      1,
				Width:     500,
				Height:    300,
				Retina:    true,
				Format:    "png256",
			},
			URL: "http://api.tiles.mapbox.com/v4/" + mapId + "/11.541650,48.151313,1/500x300@2x.png256?access_token=" + accessToken,
		},
	}

	client := NewClient(accessToken)
	for _, test := range tests {
		got, err := client.Images(mapId).buildURL(test.Request)
		if err != nil {
			t.Fatalf("expeced no error, got: %v", err)
		}
		if got != test.URL {
			t.Errorf("expected %q, got: %q", test.URL, got)
		}
	}
}

func TestGetImage(t *testing.T) {
	/*
		accessToken, err := readAccessToken(t)
		if err != nil {
			t.Fail()
			return
		}
	*/
	// Borrowed from the original developer description on Mapbox
	accessToken := "pk.eyJ1IjoibWVwbGF0byIsImEiOiJDWjV4bERjIn0.7YyMzsLszn_6k60zrl8FRw"
	mapId := "examples.map-zr0njcqy"

	client := NewClient(accessToken)
	req := &ImageRequest{
		Width:     500,
		Height:    300,
		Longitude: -73.99,
		Latitude:  40.70,
		Zoom:      13,
		Format:    "png",
	}
	img, err := client.Images(mapId).Get(req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if img == nil {
		t.Fatalf("expected image data, got: %v", img)
	}

	// TODO(oe): Check if "brooklyn.png" is actually the same file as the image we have received
}
