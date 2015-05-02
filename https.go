package yttr

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
)

func SendRequest(req *http.Request) ([]byte, error) {
	conn, err := tls.Dial("tcp", req.URL.Host+":443", nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	path := req.URL.Path + "?"
	qs := req.URL.Query()
	path += "name=" + encodeURIComponent(qs.Get("name"))
	path += "&downloadOnly=" + qs.Get("downloadOnly")
	path += "&days=" + qs.Get("days")

	_, err = io.WriteString(conn, "PUT "+path+" HTTP/1.1\r\n")
	if err != nil {
		return nil, err
	}
	for k, v := range req.Header {
		_, err = io.WriteString(conn, k+": "+v[0]+"\r\n")
		if err != nil {
			return nil, err
		}
	}
	io.WriteString(conn, "Connection: close\r\n")
	io.WriteString(conn, "User-Agent: "+Name+" cli "+Version+"\r\n")
	io.WriteString(conn, "\r\n")

	_, err = io.Copy(conn, req.Body)
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(contents, []byte("\r\n"))

	return lines[len(lines)-1], nil
}
