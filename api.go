package yttr

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type API interface {
	Upload(req UploadRequest) (UploadResponse, error)
}

func NewAPI(key, secret string) API {
	return &api{
		host:    "yttr.co",
		baseUrl: "https://yttr.co/api",
		key:     key,
		secret:  secret,
	}
}

type api struct {
	host    string
	baseUrl string
	key     string
	secret  string
}

func (a *api) Upload(ureq UploadRequest) (UploadResponse, error) {
	signature := NewSignature(a.secret, "file-upload", ureq)

	u := a.baseUrl + ureq.URLPath() + "?"

	for k, v := range *ureq.QueryString() {
		u += k + "=" + encodeURIComponent(v[0]) + "&"
	}

	req, err := http.NewRequest(ureq.HTTPMethod(), u[:len(u)-1], ureq.File())
	if err != nil {
		return nil, err
	}

	req.Header = *ureq.Header()
	req.Header.Set("Authorization", "Minister "+a.key+":"+signature)

	contents, err := SendRequest(req)
	if err != nil {
		return nil, err
	}

	rs := struct {
		Err        bool   `json:"error,omitempty"`
		Url        string `json:"url,omitempty"`
		Expiration string `json:"expiration,omitempty"`
		Kind       int    `json:"type,omitempty"`
		Message    string `json:"message,omitempty"`
	}{}

	err = json.Unmarshal(contents, &rs)
	if err != nil {
		return nil, err
	}

	if rs.Err {
		return nil, NewError(rs.Kind, rs.Message)
	}

	ru, err := url.Parse(rs.Url)
	if err != nil {
		return nil, err
	}

	d, err := parseDate(rs.Expiration)
	if err != nil {
		return nil, err
	}

	return NewUploadResponse(ru, d), nil
}
