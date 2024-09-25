package item

import (
	"github.com/gin-gonic/gin"
	"github.com/hokkung/go-groceries/internal/handler"
	"github.com/hokkung/go-groceries/internal/service/item"
	"net/http"
)

type ItemHandler struct {
	itemService item.ItemService
}

// NewItemHandler creates instance
func NewItemHandler(itemService item.ItemService) *ItemHandler {
	return &ItemHandler{itemService: itemService}
}

// ProvideItemHandler provides instance
func ProvideItemHandler(itemService item.ItemService) *ItemHandler {
	return NewItemHandler(itemService)
}

// Search searches and returns list of items
func (h *ItemHandler) Search(ctx *gin.Context) {
	var req SearchRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	details, err := h.itemService.Search(ctx, item.SearchRequest{
		Limit: req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	var res []SearchResponse
	for _, detail := range details {
		res = append(res, SearchResponse{
			ID:     detail.ID,
			URL:    detail.URL,
			Width:  detail.Width,
			Height: detail.Height,
		})
	}

	ctx.JSON(
		http.StatusOK,
		&handler.Response[[]SearchResponse]{
			Data: res,
		},
	)
}
