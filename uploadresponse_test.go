package yttr_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/txgruppi/yttr"
)

var (
	u, _ = url.Parse("https://yttr.co/o/eTTya3h0.pdf")
)

func TestUploadResponseUrl(t *testing.T) {
	res := yttr.NewUploadResponse(u, time.Now())
	equal(t, u, res.URL())
}

func TestUploadResponseExpiration(t *testing.T) {
	n := time.Now()
	res := yttr.NewUploadResponse(u, n)
	equal(t, n, res.Expiration())
}
