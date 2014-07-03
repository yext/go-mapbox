package mapbox

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/url"
)

var _ = log.Print

type ImagesAPI struct {
	c     *Client
	mapId string
}

func (api *ImagesAPI) Get(req *ImageRequest) (image.Image, error) {
	u, err := api.buildURL(req)
	if err != nil {
		return nil, err
	}

	res, err := api.c.getResponse(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// buildURL returns the complete URL for the request,
// including the access token specified in the Client.
func (api *ImagesAPI) buildURL(req *ImageRequest) (string, error) {
	urls := fmt.Sprintf("%s/%s", api.c.BaseURL(), url.QueryEscape(api.mapId))
	if len(req.Markers) > 0 {
		s := ""
		for i, marker := range req.Markers {
			if i > 0 {
				s += "/"
			}
			s += marker.String()
		}
		urls += s
	}
	urls += fmt.Sprintf("/%f,%f,%d", req.Longitude, req.Latitude, req.Zoom)
	urls += fmt.Sprintf("/%dx%d", req.Width, req.Height)
	if req.Retina {
		urls += "@2x"
	}
	urls += fmt.Sprintf(".%s", req.Format)

	u, err := url.Parse(urls)
	if err != nil {
		return "", err
	}

	// Add access_token as query string parameter
	q := u.Query()
	q.Set("access_token", api.c.accessToken)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

type ImageRequest struct {
	Latitude  float64
	Longitude float64
	Zoom      int
	Width     int
	Height    int
	Format    string
	Retina    bool
	Markers   []*Marker
}

type ImageResponse struct {
}
