package ctrl

import (
	"context"
	"errors"
	"net/http"

	"github.com/scutrobotlab/HR/dal/query"
)

var q = query.Q
var ctx = context.Background()

var ErrBadRequest error = errors.New("ErrFoo 400")
var ErrForbidden = errors.New("401")
var ErrNotFound error = errors.New("404")
var ErrInternalServer error = errors.New("500")

func errorHandle(err error) int {
	if errors.Is(err, ErrBadRequest) {
		return http.StatusBadRequest
	}
	if errors.Is(err, ErrForbidden) {
		return http.StatusForbidden
	}
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, ErrInternalServer) {
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
