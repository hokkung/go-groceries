package item

import (
	"context"
	catapi "github.com/hokkung/go-groceries/internal/client/cat_api"
)

type ItemService interface {
	Search(ctx context.Context, req SearchRequest) ([]SearchResponse, error)
}

type Item struct {
	catApi catapi.Client
}

// NewItem creates instance
func NewItem(catApi catapi.Client) *Item {
	return &Item{
		catApi: catApi,
	}
}

// ProvideItem provides instance for di
func ProvideItem(catApi catapi.Client) *Item {
	return NewItem(catApi)
}

func (i *Item) Search(ctx context.Context, req SearchRequest) ([]SearchResponse, error) {
	res, err := i.catApi.Search(ctx, catapi.SearchRequest{
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}

	var details []SearchResponse
	for _, item := range res {
		details = append(details, SearchResponse{
			ID:     item.ID,
			URL:    item.URL,
			Width:  item.Width,
			Height: item.Height,
		})
	}

	return details, nil
}
