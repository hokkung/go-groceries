package clock_test

import (
	"github.com/hokkung/go-groceries/utils/clock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var mockInstance = func() time.Time {
	return time.Date(2024, 01, 01, 10, 0, 0, 0, time.UTC)
}

func TestMain(m *testing.M) {
	clock.Freeze(mockInstance)
	m.Run()
	clock.UnFreeze()
}

func TestNow(t *testing.T) {
	exp := mockInstance()
	act := clock.Now()

	assert.Equal(t, exp, act)
}

func TestNowBKK(t *testing.T) {
	loc, _ := time.LoadLocation(clock.BangkokTZ)
	exp := mockInstance().In(loc)

	act := clock.NowBKK()

	assert.Equal(t, exp, act)
}

func TestToString(t *testing.T) {
	tt := mockInstance()
	exp := "2024-01-01 10:00:00"

	act := clock.ToString(tt)

	assert.Equal(t, exp, act)
}

func TestToDateString(t *testing.T) {
	tt := mockInstance()
	exp := "2024-01-01"

	act := clock.ToDateString(tt)

	assert.Equal(t, exp, act)
}

func TestUnFreeze(t *testing.T) {
	exp := mockInstance()

	act := clock.Now()

	assert.Equal(t, exp, act)

	clock.UnFreeze()

	act2 := clock.Now()

	assert.True(t, act.Before(act2))
}
