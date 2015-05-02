package yttr

import (
	"net/url"
	"time"
)

type UploadResponse interface {
	URL() *url.URL
	Expiration() time.Time
}

func NewUploadResponse(url *url.URL, expiration time.Time) UploadResponse {
	return &uploadResponse{
		url:        url,
		expiration: expiration,
	}
}

type uploadResponse struct {
	url        *url.URL
	expiration time.Time
}

func (r *uploadResponse) URL() *url.URL {
	return r.url
}

func (r *uploadResponse) Expiration() time.Time {
	return r.expiration
}
