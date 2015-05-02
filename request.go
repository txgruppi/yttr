package yttr

import (
	"net/http"
	"net/url"
)

type Request interface {
	HTTPMethod() string
	URLPath() string
	Header() *http.Header
	QueryString() *url.Values
	Type() string
}
