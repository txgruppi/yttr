package yttr

import (
	"net/http"
	"net/url"
)

type UploadRequest interface {
	Request
	File() File
}

func NewUploadRequest(file File) UploadRequest {
	now := formattedDate()

	header := http.Header{}
	header.Set("Content-Length", file.Size().String())
	header.Set("Content-Type", file.Type())
	header.Set("Date", now)

	queryString := url.Values{}
	queryString.Set("name", file.Name())
	queryString.Set("downloadOnly", file.DownloadOnly().String())
	queryString.Set("days", file.Days().String())

	return &uploadRequest{
		header:      &header,
		queryString: &queryString,
		file:        file,
	}
}

type uploadRequest struct {
	header      *http.Header
	queryString *url.Values
	file        File
}

func (r *uploadRequest) Header() *http.Header {
	return r.header
}

func (r *uploadRequest) QueryString() *url.Values {
	return r.queryString
}

func (r *uploadRequest) File() File {
	return r.file
}

func (r *uploadRequest) Type() string {
	return "file-upload"
}

func (r *uploadRequest) HTTPMethod() string {
	return "PUT"
}

func (r *uploadRequest) URLPath() string {
	return "/v1c/file"
}
