package mapbox

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

var _ = log.Print

type GeocodingAPI struct {
	c *Client
}

// QueryByAddress performs a forward geo-encoding. It takes an
// address and returns all kinds of information about it, e.g.
// the latitude and longitude. There might be more than one
// such result. See QueryByAddressResponse for details.
func (api *GeocodingAPI) QueryByAddress(address *QueryByAddressRequest) (*QueryByAddressResponse, error) {
	u, err := api.buildURL(address)
	if err != nil {
		return nil, err
	}

	res := new(QueryByAddressResponse)
	if err := api.c.getJSON(u, res); err != nil {
		return nil, err
	}
	return res, nil
}

// buildURL returns the complete URL for the forward geocoding request,
// including the access token specified in the Client.
func (api *GeocodingAPI) buildURL(req *QueryByAddressRequest) (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/geocoding/v5/%s/%s.json",
		api.c.BaseURL(),
		url.QueryEscape(req.Index),
		url.QueryEscape(req.Query)))
	if err != nil {
		return "", err
	}

	// Add access_token as query string parameter
	q := u.Query()
	q.Set("access_token", api.c.accessToken)

	if req.Proximity != nil {
		q.Set("proximity", fmt.Sprintf("%f,%f", req.Proximity.Latitude, req.Proximity.Longitude))
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type QueryByAddressRequest struct {
	// Index is the geocoding index as specified on
	// https://www.mapbox.com/developers/api/geocoding/.
	// If you construct a QueryByAddressRequest via
	// NewQueryByAddressRequest, the the Index is set to
	// "mapbox.places" by default.
	Index string
	// Query is the address you want to decode.
	Query string
	// Proximity is a set of coordinates
	Proximity *Coordinate
}

// NewQueryByAddressRequest creates a new QueryByAddressRequest.
// It initializes the geocoding index with "mapbox.places-v1".
func NewQueryByAddressRequest() *QueryByAddressRequest {
	return &QueryByAddressRequest{
		Index: "mapbox.places",
	}
}

// QueryByAddressResponse is the result of a call to QueryByAddress.
type QueryByAddressResponse struct {
	Attribution string                           `json:"attribution,omitempty"`
	Features    []*QueryByAddressResponseFeature `json:"features,omitempty"`
	Query       []string                         `json:"query,omitempty"`
	Type        string                           `json:"type,omitempty"`
}

type QueryByAddressResponseFeature struct {
	Type        string                 `json:"type,omitempty"`
	Text        string                 `json:"text,omitempty"`
	Relevance   float64                `json:"relevance,omitempty"`
	PlaceName   string                 `json:"place_name,omitempty"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	ID          string                 `json:"id,omitempty"`
	Geometry    *json.RawMessage       `json:"geometry,omitempty"`
	Context     *json.RawMessage       `json:"context,omitempty"`
	Center      []float64              `json:"center,omitempty"`
	BoundingBox []float64              `json:"bbox,omitempty"`
}

/*
func (api *GeocodingAPI) QueryByLonLat(lonLat *QueryByLonLatRequest) (*QueryByLonLatResponse, error) {
	return nil, nil
}

type QueryByLonLatRequest struct {
}

type QueryByLonLatResponse struct {
	Results []*QueryByLonLatResult
}

type QueryByLonLatResult struct {
}
*/
