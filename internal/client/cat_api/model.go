package catapi

import "fmt"

type SearchResponse struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type SearchRequest struct {
	Limit int
}

func (r *SearchRequest) toMapValue() map[string]string {
	return map[string]string{
		"limit": fmt.Sprintf("%d", r.Limit),
	}
}
