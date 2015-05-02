package yttr_test

import (
	"regexp"
	"testing"

	"github.com/txgruppi/yttr"
)

var (
	file   yttr.File          = nil
	req    yttr.UploadRequest = nil
	dateRe *regexp.Regexp     = nil
)

func TestUploadRequestHeader(t *testing.T) {
	equal(t, file.Size().String(), req.Header().Get("Content-Length"))
	equal(t, file.Type(), req.Header().Get("Content-Type"))
	equal(t, true, dateRe.MatchString(req.Header().Get("Date")))
}

func TestUploadRequestQueryString(t *testing.T) {
	equal(t, file.Name(), req.QueryString().Get("name"))
	equal(t, file.DownloadOnly().String(), req.QueryString().Get("downloadOnly"))
	equal(t, file.Days().String(), req.QueryString().Get("days"))
}

func TestUploadRequestFile(t *testing.T) {
	equal(t, file, req.File())
}

func TestUploadRequestType(t *testing.T) {
	equal(t, "file-upload", req.Type())
}

func TestUploadRequestHTTPMethod(t *testing.T) {
	equal(t, "PUT", req.HTTPMethod())
}

func TestUploadRequestURLPath(t *testing.T) {
	equal(t, "/v1c/file", req.URLPath())
}

func init() {
	file = yttr.NewFile(nil, "testin.txt", "text/plain", 1, 4, true)
	req = yttr.NewUploadRequest(file)
	dateRe = regexp.MustCompile("^[A-Z][a-z]{2}, [0-9]{1,2} [A-Z][a-z]{2} [0-9]{4} [0-9]{2}:[0-9]{2}:[0-9]{2} GMT$")
}
