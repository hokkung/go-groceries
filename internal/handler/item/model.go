package item

type SearchRequest struct {
	Limit int `form:"limit"`
}

type SearchResponse struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
