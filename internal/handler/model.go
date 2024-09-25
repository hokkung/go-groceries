package handler

// Response
// @Description Response Body
type Response[T any] struct {
	Data T `json:"data"`
}
