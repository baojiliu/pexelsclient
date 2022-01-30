package client

import (
	"errors"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

var (
	ErrBuildParams = errors.New("build params failed")
)

const (
	baseURL = "https://api.pexels.com/v1/"
)

type Client interface {
	SearchPhotos(req *SearchPhotosRequest) (*SearchPhotosResponse, error)
}

type httpClient struct {
	client *resty.Client
}

func (h *httpClient) SearchPhotos(req *SearchPhotosRequest) (res *SearchPhotosResponse, err error) {

	res = new(SearchPhotosResponse)

	params, err := buildParamsFromStruct(req)
	if err != nil {
		return
	}
	_, err = h.client.R().
		SetQueryParamsFromValues(params).
		SetResult(res).
		Get("search")

	if err != nil {
		return
	}

	return
}

func New(token string) Client {

	// Create a Resty Client
	client := resty.New().
		SetHostURL(baseURL).
		SetHeader("Authorization", token)

	return &httpClient{client: client}
}

func buildParamsFromStruct(v interface{}) (url.Values, error) {
	params, err := query.Values(v)
	if err != nil {
		return nil, ErrBuildParams
	}
	return params, nil
}
