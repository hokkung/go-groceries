package item_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	catapi "github.com/hokkung/go-groceries/internal/client/cat_api"
	mock_catapi "github.com/hokkung/go-groceries/internal/client/cat_api/mock"
	"github.com/hokkung/go-groceries/internal/service/item"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ItemServiceSuite struct {
	suite.Suite

	ctx        context.Context
	mockErr    error
	mockCatApi *mock_catapi.MockClient

	underTest *item.Item
}

func (ts *ItemServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(ts.T())

	ts.ctx = context.Background()
	ts.mockErr = errors.New("mock")
	ts.mockCatApi = mock_catapi.NewMockClient(ctrl)
	ts.underTest = item.ProvideItem(ts.mockCatApi)
}

func TestItemServiceSuite(t *testing.T) {
	suite.Run(t, new(ItemServiceSuite))
}

func (ts *ItemServiceSuite) TestSearch() {
	ts.Run("should be returned error with Search", func() {
		mockReq := item.SearchRequest{
			Limit: 1,
		}

		ts.mockCatApi.
			EXPECT().
			Search(ts.ctx, catapi.SearchRequest{
				Limit: mockReq.Limit,
			}).
			Return(nil, ts.mockErr)

		res, err := ts.underTest.Search(ts.ctx, mockReq)

		ts.Equal(ts.mockErr, err)
		ts.Nil(res)
	})

	ts.Run("should be returned ok", func() {
		mockReq := item.SearchRequest{
			Limit: 1,
		}
		mockRes := []catapi.SearchResponse{
			{
				ID:     "1",
				URL:    "url",
				Width:  100,
				Height: 100,
			},
		}

		ts.mockCatApi.
			EXPECT().
			Search(ts.ctx, catapi.SearchRequest{
				Limit: mockReq.Limit,
			}).
			Return(mockRes, nil)

		res, err := ts.underTest.Search(ts.ctx, mockReq)

		ts.Nil(err)
		ts.Equal(1, len(res))
		for i, val := range mockRes {
			ts.Equal(val.ID, res[i].ID)
			ts.Equal(val.URL, res[i].URL)
			ts.Equal(val.Width, res[i].Width)
			ts.Equal(val.Height, res[i].Height)
		}
	})
}
