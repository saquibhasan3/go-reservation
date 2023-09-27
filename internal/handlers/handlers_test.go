package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	// {"contact-thankyou", "/contact-thankyou", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "POST", []postData{
		{key: "start", value: "2023-05-01"},
		{key: "end", value: "2023-05-01"},
		{key: "adults", value: "5"},
		{key: "children", value: "2"},
	}, http.StatusOK},
	{"search-availability", "/search-availability", "POST", []postData{
		{key: "first_name", value: "Saquib"},
		{key: "last_name", value: "Hasaan"},
		{key: "email", value: "saquibhasan3@gmail.com"},
		{key: "phone", value: "9520115711"},
		{key: "message", value: "Any message which user send."},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
