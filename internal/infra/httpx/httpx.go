// internal/httpx/httpx.go
package httpx

import (
	"net/http"
	"rest-go/internal/shared/errs"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Wrap(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
			return
		}

		if httpErr, ok := err.(errs.HttpErrorInterface); ok {
			errs.Throw(w, httpErr, r)
			return
		}

		errs.Throw(w, errs.NewHttpError("Internal Server Error", "internal_error", http.StatusInternalServerError), r)

	}
}
