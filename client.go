package mapbox

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	MapboxHost       = "api.tiles.mapbox.com"
	MapboxPathPrefix = "/v4"
	UserAgent        = "Mapbox Google Go Client v0.1"
)

type Client struct {
	httpClient  *http.Client
	https       bool
	accessToken string
}

func NewClient(accessToken string) *Client {
	return &Client{
		https:       false,
		accessToken: accessToken,
		httpClient:  http.DefaultClient,
	}
}

func (c *Client) SetHTTPClient(client *http.Client) {
	if client == nil {
		client = http.DefaultClient
	}
	c.httpClient = client
}

func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

func (c *Client) SetHTTPS(https bool) {
	c.https = https
}

func (c *Client) HTTPS() bool {
	return c.https
}

func (c *Client) BaseURL() string {
	url := new(url.URL)
	if c.https {
		url.Scheme = "https"
	} else {
		url.Scheme = "http"
	}
	url.Host = MapboxHost
	url.Path = MapboxPathPrefix
	return url.String()
}

func (c *Client) Images(mapId string) *ImagesAPI {
	return &ImagesAPI{c: c, mapId: mapId}
}

func (c *Client) Geocoding() *GeocodingAPI {
	return &GeocodingAPI{c: c}
}

// getResponse returns the HTTP response to the caller.
// Warning: The caller is responsible for closing the
// Body via e.g. `defer res.Body.Close()`.
func (c *Client) getResponse(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	req.Header.Set("User-Agent", UserAgent)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) getJSON(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", UserAgent)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}
