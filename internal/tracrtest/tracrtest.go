package tracrtest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
)

//nb: Following the httptest approach, for test utils
// https://cs.opensource.google/go/go/+/master:src/net/http/httptest/httptest.go

// MockHttpRequest - helper, to build a request in tests
//
//	e.g. 	req := tracrtest.MockHttpRequest("GET", "/example/", nil, url.Values{
//			"blah": {"blahblah"},
//		}, header)
//
// nb: Note the url.Values extra braces, as a key can have multiple values
func MockHttpRequest(method, url string, body io.Reader, params url.Values, header http.Header) *http.Request {
	req := httptest.NewRequest(method, url, body)

	// add params
	req.URL.RawQuery = params.Encode()

	if header != nil {
		req.Header = header
	}
	return req
}
