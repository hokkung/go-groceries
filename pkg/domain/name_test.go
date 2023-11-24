package domain_test

import (
	"testing"

	"github.com/hokkung/go-groceries/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestNameModel(t *testing.T) {
	d1 := domain.Name{
		Primary:     "pri",
		ThaiName:    "th",
		EnglishName: "eng",
	}

	assert.Equal(t, d1.Primary, "pri")
	assert.Equal(t, d1.ThaiName, "th")
	assert.Equal(t, d1.EnglishName, "eng")
}
