package yttr

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func NewSignature(secret, kind string, req Request) string {
	str := buildSignatureString(kind, req)
	str = hashSignatureString(secret, str)
	return str
}

func hashSignatureString(secret, signature string) string {
	hm := hmac.New(sha1.New, []byte(secret))
	hm.Write([]byte(signature))
	return fmt.Sprintf("%x", hm.Sum(nil))
}

func buildSignatureString(kind string, req Request) string {
	var b bytes.Buffer

	b.WriteString("HEADERS\n")
	b.WriteString("content-length: " + req.Header().Get("Content-Length") + "\n")
	b.WriteString("content-type: " + req.Header().Get("Content-Type") + "\n")
	b.WriteString("date: " + req.Header().Get("Date") + "\n\n")
	b.WriteString("QUERY\n")
	b.WriteString("name: " + req.QueryString().Get("name") + "\n")
	b.WriteString("downloadOnly: " + req.QueryString().Get("downloadOnly") + "\n")
	b.WriteString("days: " + req.QueryString().Get("days") + "\n\n")
	b.WriteString("TYPE\n")
	b.WriteString(kind + "\n\n")

	return b.String()
}
