package clock

import (
	"sync"
	"time"
)

const (
	BangkokTZ = "Asia/Bangkok"

	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
)

var (
	instance = time.Now
	lock     = sync.Mutex{}
)

// NowBKK returns current time in BKK timezone
func NowBKK() time.Time {
	loc, _ := time.LoadLocation(BangkokTZ)
	return Now().In(loc)
}

// Now returns current time
func Now() time.Time {
	return instance()
}

// ToString returns date time in string format
func ToString(n time.Time) string {
	return n.Format(DateTimeFormat)
}

// ToDateString returns date in string format
func ToDateString(n time.Time) string {
	return n.Format(DateFormat)
}

// Freeze sets the instance of time
func Freeze(fn func() time.Time) {
	lock.Lock()
	defer lock.Unlock()
	instance = fn
}

// UnFreeze removes all applied time
func UnFreeze() {
	lock.Lock()
	defer lock.Unlock()

	instance = time.Now
}
