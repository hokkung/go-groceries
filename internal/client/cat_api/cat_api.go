package catapi

import (
	"context"
	"errors"
	"github.com/hokkung/go-groceries/pkg/cnt"
)

//go:generate mockgen -source ./cat_api.go -destination ./mock/mock_cat_api.go
type Client interface {
	Search(ctx context.Context, req SearchRequest) ([]SearchResponse, error)
}

type CatAPI struct {
	client *cnt.Client
}

// NewCatAPI creates instance
func NewCatAPI() (*CatAPI, error) {
	client, err := cnt.NewClient(&cnt.ClientOptions{
		Name:    "cat-api",
		BaseURL: "https://api.thecatapi.com",
	})
	if err != nil {
		return nil, err
	}

	return &CatAPI{
		client: client,
	}, nil
}

// ProvideCatAPI provides instance for di
func ProvideCatAPI() (*CatAPI, error) {
	return NewCatAPI()
}

// Search searches
func (c CatAPI) Search(
	ctx context.Context,
	req SearchRequest,
) ([]SearchResponse, error) {
	var response []SearchResponse
	res, err := c.client.
		Req(ctx).
		SetQueryParams(req.toMapValue()).
		SetResult(&response).
		Get(SearchURL)
	if err != nil {
		return nil, err
	}

	if !res.IsSuccess() {
		return nil, errors.New("search cat failed")
	}

	return response, nil
}
